package services

import (
	"database/sql"
	"finance-tracker/models"
)

func GetAllAccounts(db *sql.DB, userID int) ([]models.Account, error) {
	rows, err := db.Query("SELECT id, name, type, balance FROM accounts WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		var acc models.Account
		if err := rows.Scan(&acc.ID, &acc.Name, &acc.Type, &acc.Balance); err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}
	return accounts, nil
}
