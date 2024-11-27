package models

import "time"

type Expense struct {
	ID          int       `json:"id"`
	Category    string    `json:"category"`
	Date        string    `json:"date"`
	Description string    `json:"description"`
	Amount      float32   `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
