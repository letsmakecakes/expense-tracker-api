package models

import "time"

// ExpenseCategory represents the category of an expense.
type ExpenseCategory string

// List of expense categories.
const (
	CategoryGroceries   ExpenseCategory = "Groceries"
	CategoryLeisure     ExpenseCategory = "Leisure"
	CategoryElectronics ExpenseCategory = "Electronics"
	CategoryUtilities   ExpenseCategory = "Utilities"
	CategoryClothing    ExpenseCategory = "Clothing"
	CategoryHealth      ExpenseCategory = "Health"
	CategoryOthers      ExpenseCategory = "Others"
)

// Expense represents an expense record with relevant details.
type Expense struct {
	ID          int             `json:"id"`          // Unique identifier for the expense
	UserID      int             `json:"user_id"`     // Unique identifier for the user who made the expense
	Amount      float64         `json:"amount"`      // Amount of the expense
	Category    ExpenseCategory `json:"category"`    // Category of the expense
	Description string          `json:"description"` // Description of the expense
	Date        time.Time       `json:"date"`        // Date of the expense
	CreatedAt   time.Time       `json:"created_at"`  // Timestamp when the record was created
	UpdatedAt   time.Time       `json:"updated_at"`  // Timestamp when the record was last updated
}
