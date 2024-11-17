package repository

import (
	"database/sql"
	"expensetrackerapi/pkg/models"
)

type CredentialRepository interface {
	Create(credential *models.Credential) error
	GetByUsername(username string) (*models.Credential, error)
	Update(credential *models.Credential) error
	Delete(id int) error
}

type credentialRepository struct {
	db *sql.DB
}

func NewCredentialRepository(db *sql.DB) CredentialRepository {
	return &credentialRepository{db}
}

func (r *credentialRepository) Create(credential *models.Credential) error {
	query := `INSERT INTO login (username, password, created_at, updated_at) 
				VALUES ($1, $2, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, credential.Username, credential.Password).Scan(&credential.ID, &credential.CreatedAt, &credential.UpdatedAt)
	return err
}

func (r *credentialRepository) GetByUsername(username string) (*models.Credential, error) {
	query := `SELECT id, username, password, created_at, updated_at FROM login WHERE username = $1`
	row := r.db.QueryRow(query, username)

	var credential models.Credential
	err := row.Scan(&credential.ID, &credential.Username, &credential.Password, &credential.CreatedAt, &credential.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &credential, nil
}

func (r *credentialRepository) Update(credential *models.Credential) error {
	query := `UPDATE login SET username = $1, password = $2, updated_at = NOW() WHERE id = $3 RETURNING updated_at`
	err := r.db.QueryRow(query, credential.Username, credential.Password, credential.ID).Scan(&credential.UpdatedAt)
	return err
}

func (r *credentialRepository) Delete(id int) error {
	query := `DELETE FROM login WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
