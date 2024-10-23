package repository

import (
	"database/sql"
	"expensetrackerapi/pkg/models"
	"strings"
)

type CredentialRepository interface {
	Create(credential *models.Credentials) error
	GetByID(id int) (*models.Credentials, error)
	Update(credential *models.Credentials) error
	Delete(id int) error
}

type credentialRepository struct {
	db *sql.DB
}

func NewCredentialRepository(db *sql.DB) CredentialRepository {
	return &credentialRepository{db}
}

func (r *credentialRepository) Create(credential *models.Credentials) error {
	query := `INSERT INTO user (username, password, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, credential.Username, credential.Password, credential.CreatedAt, credential.UpdatedAt).Scan(&credential.ID, &credential.CreatedAt, &credential.UpdatedAt)
	return err
}
