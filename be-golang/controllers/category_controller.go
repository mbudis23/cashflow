package controllers

import (
	"database/sql"
	"finance-tracker/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GET /categories
func GetCategories(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id") // diasumsikan dari JWT middleware

		rows, err := db.Query("SELECT id, name, color FROM categories WHERE user_id = $1", userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var categories []models.Category
		for rows.Next() {
			var cat models.Category
			if err := rows.Scan(&cat.ID, &cat.Name, &cat.Color); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			categories = append(categories, cat)
		}
		c.JSON(http.StatusOK, categories)
	}
}

// POST /categories
func CreateCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")

		var cat models.Category
		if err := c.ShouldBindJSON(&cat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `INSERT INTO categories (user_id, name, color) VALUES ($1, $2, $3) RETURNING id`
		err := db.QueryRow(query, userID, cat.Name, cat.Color).Scan(&cat.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, cat)
	}
}

// PUT /categories/:id
func UpdateCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryID := c.Param("id")

		var cat models.Category
		if err := c.ShouldBindJSON(&cat); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		query := `UPDATE categories SET name=$1, color=$2 WHERE id=$3`
		_, err := db.Exec(query, cat.Name, cat.Color, categoryID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Category updated"})
	}
}

// DELETE /categories/:id
func DeleteCategory(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryID := c.Param("id")

		id, err := strconv.Atoi(categoryID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
			return
		}

		_, err = db.Exec("DELETE FROM categories WHERE id=$1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
	}
}
