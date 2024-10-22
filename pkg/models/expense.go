package models

import "time"

type Expense struct {
	ID          int       `json:"id"`
	Date        string    `json:"date"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
