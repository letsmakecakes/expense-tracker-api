package repository

import "expensetrackerapi/pkg/models"

type ExpenseRepository interface {
	Add(expense *models.Expense) error
	Load() ([]*models.Expense, error)
	Update(expense *models.Expense) error
	Delete(id int) error
}
