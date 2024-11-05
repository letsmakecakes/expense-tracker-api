package repository

import (
	"database/sql"
	"expensetrackerapi/pkg/models"
)

type ExpenseRepository interface {
	Add(expense *models.Expense) error
	GetByID(id int) (*models.Expense, error)
	LoadAll() ([]*models.Expense, error)
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
	query := `INSERT INTO expense (date, description, amount, created_at, updated_at)
				VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, expense.Date, expense.Description, expense.Amount).Scan(&expense.ID, &expense.CreatedAt, &expense.UpdatedAt)
	return err
}

func (r *expenseRepository) GetByID(id int) (*models.Expense, error) {
	query := `SELECT id, date, description, amount, created_at, updated_at FROM expense WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var expense models.Expense
	err := row.Scan(&expense.ID, &expense.Date, &expense.Description, &expense.Amount, &expense.CreatedAt, &expense.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func (r *expenseRepository) LoadAll() ([]*models.Expense, error) {
	var expenses []*models.Expense
	var rows *sql.Rows
	var err error

	query := `SELECT id, date, description, amount, created_at, updated_at FROM expense`
	rows, err = r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(&expense.ID, &expense.Date, &expense.Description, &expense.Amount, &expense.CreatedAt, &expense.UpdatedAt); err != nil {
			return nil, err
		}
		expenses = append(expenses, &expense)
	}

	return expenses, nil
}

func (r *expenseRepository) Update(expense *models.Expense) error {
	query := `UPDATE expense SET date = $1, description = $2, amount = $3, updated_at = NOW() WHERE id = $4 RETURNING updated_at`
	err := r.db.QueryRow(query, expense.Date, expense.Description, expense.Amount, expense.ID).Scan(&expense.UpdatedAt)
	return err
}

func (r *expenseRepository) Delete(id int) error {
	query := `DELETE FROM expense WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
