package services

import (
	"expensetrackerapi/internal/repository"
	"expensetrackerapi/pkg/models"
)

type ExpenseService interface {
	AddExpense(expense *models.Expense) error
	GetExpenseByID(id int) (*models.Expense, error)
	LoadAllExpenses() ([]*models.Expense, error)
	UpdateExpense(expense *models.Expense) error
	DeleteExpense(id int) error
}

type expenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseRepository(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{repo}
}

func (s *expenseService) AddExpense(expense *models.Expense) error {
	return s.repo.Add(expense)
}

func (s *expenseService) GetExpenseByID(id int) (*models.Expense, error) {
	return s.repo.GetByID(id)
}

func (s *expenseService) LoadAllExpenses() ([]*models.Expense, error) {
	return s.repo.LoadAll()
}