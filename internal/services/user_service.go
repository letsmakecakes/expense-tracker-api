package services

import (
	"expensetrackerapi/internal/repository"
	"expensetrackerapi/pkg/models"
)

type CredentialService interface {
	CreateCredential(credential *models.Credentials) error
	GetCredentialByID(credential *models.Credentials) error
	UpdateCredential(credential *models.Credentials) error
	DeleteCredential(id int) error
}

type credentialService struct {
	repo repository.CredentialRepository
}

func NewCredentialService(repo repository.CredentialRepository) CredentialService {
	return &credentialService{repo}
}

func (s *credentialService) CreateCredential(credential *models.Credentials) error {
	return s.repo.Create(credential)
}
