package services

import (
	"expensetrackerapi/internal/repository"
	"expensetrackerapi/pkg/models"
)

type CredentialService interface {
	CreateCredential(credential *models.Credential) error
	GetCredentialByUsername(username string) (*models.Credential, error)
	UpdateCredential(credential *models.Credential) error
	DeleteCredential(id int) error
}

type credentialService struct {
	repo repository.CredentialRepository
}

func NewCredentialService(repo repository.CredentialRepository) CredentialService {
	return &credentialService{repo}
}

func (s *credentialService) CreateCredential(credential *models.Credential) error {
	return s.repo.Create(credential)
}

func (s *credentialService) GetCredentialByUsername(username string) (*models.Credential, error) {
	return s.repo.GetByUsername(username)
}

func (s *credentialService) UpdateCredential(credential *models.Credential) error {
	return s.repo.Update(credential)
}

func (s *credentialService) DeleteCredential(id int) error {
	return s.repo.Delete(id)
}
