package services

import (
	"database/sql"
	"finance-tracker/models"
)

func GetAllTransactions(db *sql.DB, userID int) ([]models.Transaction, error) {
	rows, err := db.Query(`SELECT id, type, amount, account_from, account_to, category_id, description, is_tax_applied, tax_amount, note, date
	                       FROM transactions WHERE user_id = $1 ORDER BY date DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		if err := rows.Scan(&t.ID, &t.Type, &t.Amount, &t.AccountFrom, &t.AccountTo, &t.CategoryID,
			&t.Description, &t.IsTaxApplied, &t.TaxAmount, &t.Note, &t.Date); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
