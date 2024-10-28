package controllers

import "expensetrackerapi/internal/services"

type CredentialController struct {
	Service services.CredentialService
}

func NewCredentialController(service services.CredentialService) *CredentialController {
	return &CredentialController{service}
}

