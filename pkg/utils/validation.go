package utils

import (
	"errors"
	"expensetrackerapi/pkg/models"
	"strings"
)

func ValidateExpense(expense *models.Expense) error {
	if strings.TrimSpace(expense.Date) == "" {
		return errors.New("Date is required")
	}
	if strings.TrimSpace(expense.Description) == "" {
		return errors.New("Description is required")
	}
	if expense.Amount == 0 {
		return errors.New("Amount is required")
	}

	return nil
}
