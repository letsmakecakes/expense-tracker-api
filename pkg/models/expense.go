package models

import "time"

// Expense represents an expense record with relevant details.
type Expense struct {
	ID          int       `json:"id"`          // Unique identifier for the expense
	UserID      int       `json:"user_id"`     // Unique identifier for the user who made the expense
	Amount      float64   `json:"amount"`      // Amount of the expense
	Category    string    `json:"category"`    // Category of the expense
	Description string    `json:"description"` // Description of the expense
	Date        string    `json:"date"`        // Date of the expense
	CreatedAt   time.Time `json:"created_at"`  // Timestamp when the record was created
	UpdatedAt   time.Time `json:"updated_at"`  // Timestamp when the record was last updated
}
