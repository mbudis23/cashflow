package services

import (
	"database/sql"
	"finance-tracker/models"
)

func GetAllLoans(db *sql.DB, userID int) ([]models.Loan, error) {
	rows, err := db.Query(`SELECT id, contact_name, amount, loan_type, due_date, interest, is_paid, note, related_transaction_id
	                       FROM loans WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var loans []models.Loan
	for rows.Next() {
		var l models.Loan
		if err := rows.Scan(&l.ID, &l.ContactName, &l.Amount, &l.LoanType, &l.DueDate,
			&l.Interest, &l.IsPaid, &l.Note, &l.RelatedTransactionID); err != nil {
			return nil, err
		}
		loans = append(loans, l)
	}
	return loans, nil
}
