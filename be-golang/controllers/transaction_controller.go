package controllers

import (
	"database/sql"
	"finance-tracker/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /transactions
func GetTransactions(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		query := `SELECT id, type, amount, account_from, account_to, category_id, description, is_tax_applied, tax_amount, note, date
		          FROM transactions WHERE user_id = $1 ORDER BY date DESC`

		rows, err := db.Query(query, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var transactions []models.Transaction
		for rows.Next() {
			var t models.Transaction
			err := rows.Scan(&t.ID, &t.Type, &t.Amount, &t.AccountFrom, &t.AccountTo, &t.CategoryID,
				&t.Description, &t.IsTaxApplied, &t.TaxAmount, &t.Note, &t.Date)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			transactions = append(transactions, t)
		}
		c.JSON(http.StatusOK, transactions)
	}
}

// POST /transactions
func CreateTransaction(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		var t models.Transaction
		if err := c.ShouldBindJSON(&t); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `INSERT INTO transactions (user_id, type, amount, account_from, account_to, category_id,
		          description, is_tax_applied, tax_amount, note, date)
		          VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`

		err := db.QueryRow(query, userID, t.Type, t.Amount, t.AccountFrom, t.AccountTo,
			t.CategoryID, t.Description, t.IsTaxApplied, t.TaxAmount, t.Note, t.Date).Scan(&t.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, t)
	}
}

// PUT /transactions/:id
func UpdateTransaction(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionID := c.Param("id")

		var t models.Transaction
		if err := c.ShouldBindJSON(&t); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `UPDATE transactions SET type=$1, amount=$2, account_from=$3, account_to=$4, category_id=$5,
		          description=$6, is_tax_applied=$7, tax_amount=$8, note=$9, date=$10 WHERE id=$11`

		_, err := db.Exec(query, t.Type, t.Amount, t.AccountFrom, t.AccountTo, t.CategoryID,
			t.Description, t.IsTaxApplied, t.TaxAmount, t.Note, t.Date, transactionID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Transaction updated"})
	}
}

// DELETE /transactions/:id
func DeleteTransaction(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionID := c.Param("id")

		id, err := strconv.Atoi(transactionID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		_, err = db.Exec("DELETE FROM transactions WHERE id=$1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
	}
}
