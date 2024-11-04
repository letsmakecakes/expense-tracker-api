package controllers

import (
	"database/sql"
	"errors"
	"expensetrackerapi/internal/services"
	"expensetrackerapi/pkg/jwt"
	"expensetrackerapi/pkg/models"
	"expensetrackerapi/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CredentialController struct {
	Service services.CredentialService
}

func NewCredentialController(service services.CredentialService) *CredentialController {
	return &CredentialController{service}
}

// CreateCredential handles POST /user/signup
func (c *CredentialController) CreateCredential(ctx *gin.Context) {
	var credential models.Credential
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		log.Errorf("error binding JSON: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the credential
	if err := utils.ValidateCredential(&credential); err != nil {
		log.Errorf("error validating blog: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Check if the user already exists
	existingCred, err := c.Service.GetCredentialByID(credential.ID)
	if err != nil {
		log.Errorf("error getting credential: %v", err)
		if err != sql.ErrNoRows {
			utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}

	if existingCred != nil {
		log.Errorf("credential already exists")
		utils.RespondWithError(ctx, http.StatusBadRequest, "Credential already exists")
		return
	}

	hashed_password, err := utils.HashPassword(credential.Password)
	if err != nil {
		log.Errorf("error hashing password: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	credential.Password = hashed_password

	if err := c.Service.CreateCredential(&credential); err != nil {
		log.Errorf("error creating credential: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to create credential")
		return
	}

	utils.RespondWithJSON(ctx, http.StatusCreated, credential)
}

// GetCredential handles GET /user/login
func (c *CredentialController) GetCredential(ctx *gin.Context) {
	var credential models.Credential
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		log.Errorf("error binding JSON: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	existingCredential, err := c.Service.GetCredentialByID(credential.ID)
	if err != nil {
		log.Errorf("error getting credential: %v", err)
		if err == sql.ErrNoRows {
			utils.RespondWithError(ctx, http.StatusNotFound, "Credentail not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve credential")
		}
		return
	}

	match := utils.CheckPasswordHash(credential.Password, existingCredential.Password)
	if !match {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid credentials")
		return
	}

	// Generate a JWT
	token, err := jwt.GenerateToken(credential.Username)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Could not generate token")
		return
	}

	// Set the JWT as an HTTP-only cookie
	ctx.SetCookie("token", token, 43200, "/", "localhost", false, true)

	utils.RespondWithJSON(ctx, http.StatusOK, "Logged in successfully")
}

// UpdateCredential handles PUT /user/:id
func (c *CredentialController) UpdateCredential(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("error updating credential: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid post ID")
		return
	}

	var credential models.Credential
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate the credential
	if err := utils.ValidateCredential(&credential); err != nil {
		log.Errorf("error validating credential: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	credential.ID = id
	if err := c.Service.UpdateCredential(&credential); err != nil {
		log.Errorf("error updating credential: %v", err)
		if err == sql.ErrNoRows {
			utils.RespondWithError(ctx, http.StatusNotFound, "Credential not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to update credential")
		}
		return
	}

}

// DeleteCredential handles DELETE /user/:id
func (c *CredentialController) DeleteCredential(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Errorf("error converting parameter to integer: %v", err)
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid blog ID")
		return
	}

	if err := c.Service.DeleteCredential(id); err != nil {
		log.Errorf("error deleting credential: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(ctx, http.StatusNotFound, "Credential not found")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to delete credential")
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}
