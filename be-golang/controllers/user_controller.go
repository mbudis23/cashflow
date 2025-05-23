package controllers

import (
	"database/sql"
	"finance-tracker/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser membuat user baru
func RegisterUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
			return
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
			return
		}
		user.Password = string(hashedPassword)

		query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
		err = db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Email already registered or DB error"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created", "user_id": user.ID})
	}
}

// LoginUser memverifikasi user dan memberikan JWT
func LoginUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.User
		var stored models.User

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
			return
		}

		query := `SELECT id, name, email, password FROM users WHERE email=$1`
		err := db.QueryRow(query, input.Email).Scan(&stored.ID, &stored.Name, &stored.Email, &stored.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		// Compare password
		if err := bcrypt.CompareHashAndPassword([]byte(stored.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
			return
		}

		// Generate JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": stored.ID,
			"email":   stored.Email,
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
			"user":  gin.H{"id": stored.ID, "name": stored.Name, "email": stored.Email},
		})
	}
}
