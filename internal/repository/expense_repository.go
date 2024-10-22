package repository

import (
	"database/sql"
	"expensetrackerapi/pkg/models"
)

type ExpenseRepository interface {
	Add(expense *models.Expense) error
	Load() ([]*models.Expense, error)
	Update(expense *models.Expense) error
	Delete(id int) error
}

type expenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) ExpenseRepository {
	return &expenseRepository{db}
}

func (r *expenseRepository) Add(expense *models.Expense) error {
	query := `INSERT INTO blogs (date, description, amount, created_at, updated_at)
				VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, expense.Date, expense.Description, expense.Amount).Scan(&expense.ID, &expense.CreatedAt, &expense.UpdatedAt)
	return err
}
