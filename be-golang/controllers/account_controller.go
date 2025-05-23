package controllers

import (
	"database/sql"
	"finance-tracker/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /accounts
func GetAccounts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id") // diasumsikan dari middleware JWT

		rows, err := db.Query("SELECT id, name, type, balance FROM accounts WHERE user_id = $1", userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		accounts := []models.Account{}
		for rows.Next() {
			var acc models.Account
			if err := rows.Scan(&acc.ID, &acc.Name, &acc.Type, &acc.Balance); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			accounts = append(accounts, acc)
		}
		c.JSON(http.StatusOK, accounts)
	}
}

// POST /accounts
func CreateAccount(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		var acc models.Account
		if err := c.ShouldBindJSON(&acc); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `INSERT INTO accounts (user_id, name, type, balance) VALUES ($1, $2, $3, $4) RETURNING id`
		err := db.QueryRow(query, userID, acc.Name, acc.Type, acc.Balance).Scan(&acc.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, acc)
	}
}

// PUT /accounts/:id
func UpdateAccount(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		accountID := c.Param("id")
		var acc models.Account

		if err := c.ShouldBindJSON(&acc); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `UPDATE accounts SET name=$1, type=$2, balance=$3 WHERE id=$4`
		_, err := db.Exec(query, acc.Name, acc.Type, acc.Balance, accountID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Account updated"})
	}
}

// DELETE /accounts/:id
func DeleteAccount(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		accountID := c.Param("id")

		id, err := strconv.Atoi(accountID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}

		_, err = db.Exec("DELETE FROM accounts WHERE id = $1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Account deleted"})
	}
}
