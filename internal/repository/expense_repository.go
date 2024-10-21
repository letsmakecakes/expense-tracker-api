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
