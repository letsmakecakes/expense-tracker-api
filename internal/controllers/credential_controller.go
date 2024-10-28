package controllers

import (
	"expensetrackerapi/internal/services"
	"expensetrackerapi/pkg/models"
	"expensetrackerapi/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type CredentialController struct {
	Service services.CredentialService
}

func NewCredentialController(service services.CredentialService) *CredentialController {
	return &CredentialController{service}
}

// CreateCredential handles POST /signup
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
