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
	//existingCred, err := c.Service.GetCredentialByUsername(credential.ID)
	//if err != nil {
	//	log.Errorf("error getting credential: %v", err)
	//	if !errors.Is(err, sql.ErrNoRows) {
	//		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
	//		return
	//	}
	//}
	//
	//if existingCred != nil {
	//	log.Errorf("credential already exists")
	//	utils.RespondWithError(ctx, http.StatusBadRequest, "Credential already exists")
	//	return
	//}

	hashedPassword, err := utils.HashPassword(credential.Password)
	if err != nil {
		log.Errorf("error hashing password: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	credential.Password = hashedPassword

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

	// Parse the request body into the Credential struct
	if err := ctx.ShouldBindJSON(&credential); err != nil {
		log.Errorf("error binding JSON: %v", err)

		// Optional: Log raw data if binding fails
		if rawData, errRaw := ctx.GetRawData(); errRaw == nil {
			log.Errorf("received raw data: %v", string(rawData))
		} else {
			log.Errorf("error retrieving raw data: %v", errRaw)
		}

		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate input fields (ensure non-empty values)
	if credential.Username == "" || credential.Password == "" {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Username and password are required")
		return
	}

	// Retrieve credential from the database using the service
	existingCredential, err := c.Service.GetCredentialByUsername(credential.Username)
	if err != nil {
		log.Errorf("error getting credential: %v", err)
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(ctx, http.StatusUnauthorized, "Invalid username or password")
		} else {
			utils.RespondWithError(ctx, http.StatusInternalServerError, "Failed to retrieve credential")
		}
		return
	}

	// Verify the password
	match := utils.CheckPasswordHash(credential.Password, existingCredential.Password)
	if !match {
		utils.RespondWithError(ctx, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	// Generate a JWT token
	token, err := jwt.GenerateToken(existingCredential.Username)
	if err != nil {
		log.Errorf("error generating JWT token: %v", err)
		utils.RespondWithError(ctx, http.StatusInternalServerError, "Could not generate token")
		return
	}

	// Option 1: Set the JWT as an HTTP-only cookie (if you want cookie-based authentication)
	ctx.SetCookie("token", token, 3600, "/", "localhost", false, true)

	// Option 2: Set the JWT as a response header (if token-based auth is preferred)
	ctx.Header("Authorization", "Bearer "+token)

	// Respond with a success message
	utils.RespondWithJSON(ctx, http.StatusOK, gin.H{
		"message": "Logged in successfully",
		"token":   token, // Include token in the response if needed
	})
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
		if errors.Is(err, sql.ErrNoRows) {
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
