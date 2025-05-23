package services

import (
	"database/sql"
	"finance-tracker/models"
)

func GetUserByEmail(db *sql.DB, email string) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}
