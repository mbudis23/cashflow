package controllers

import (
	"database/sql"
	"finance-tracker/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /loans
func GetLoans(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		query := `SELECT id, contact_name, amount, loan_type, due_date, interest, is_paid, note, related_transaction_id
		          FROM loans WHERE user_id = $1 ORDER BY due_date`

		rows, err := db.Query(query, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var loans []models.Loan
		for rows.Next() {
			var loan models.Loan
			err := rows.Scan(&loan.ID, &loan.ContactName, &loan.Amount, &loan.LoanType, &loan.DueDate,
				&loan.Interest, &loan.IsPaid, &loan.Note, &loan.RelatedTransactionID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			loans = append(loans, loan)
		}
		c.JSON(http.StatusOK, loans)
	}
}

// POST /loans
func CreateLoan(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		var loan models.Loan
		if err := c.ShouldBindJSON(&loan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `INSERT INTO loans (user_id, contact_name, amount, loan_type, due_date,
		          interest, is_paid, note, related_transaction_id)
		          VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`

		err := db.QueryRow(query, userID, loan.ContactName, loan.Amount, loan.LoanType,
			loan.DueDate, loan.Interest, loan.IsPaid, loan.Note, loan.RelatedTransactionID).Scan(&loan.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, loan)
	}
}

// PUT /loans/:id
func UpdateLoan(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		loanID := c.Param("id")

		var loan models.Loan
		if err := c.ShouldBindJSON(&loan); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `UPDATE loans SET contact_name=$1, amount=$2, loan_type=$3, due_date=$4,
		          interest=$5, is_paid=$6, note=$7, related_transaction_id=$8 WHERE id=$9`

		_, err := db.Exec(query, loan.ContactName, loan.Amount, loan.LoanType, loan.DueDate,
			loan.Interest, loan.IsPaid, loan.Note, loan.RelatedTransactionID, loanID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Loan updated"})
	}
}

// DELETE /loans/:id
func DeleteLoan(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		loanID := c.Param("id")

		id, err := strconv.Atoi(loanID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid loan ID"})
			return
		}

		_, err = db.Exec("DELETE FROM loans WHERE id=$1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Loan deleted"})
	}
}
