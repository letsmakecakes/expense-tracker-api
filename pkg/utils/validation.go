package utils

import (
	"errors"
	"expensetrackerapi/pkg/models"
	"strings"
)

func ValidateExpense(expense *models.Expense) error {
	if strings.TrimSpace(expense.Date) == "" {
		return errors.New("date is required")
	}
	if strings.TrimSpace(expense.Description) == "" {
		return errors.New("description is required")
	}
	if expense.Amount == 0 {
		return errors.New("amount is required")
	}

	return nil
}

func ValidateCredential(credential *models.Credential) error {
	if strings.TrimSpace(credential.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(credential.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}
