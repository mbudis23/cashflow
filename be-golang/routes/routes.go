package routes

import (
	"database/sql"
	"finance-tracker/controllers"
	"finance-tracker/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/register", controllers.RegisterUser(db))
	router.POST("/login", controllers.LoginUser(db))

	authorized := router.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.GET("/accounts", controllers.GetAccounts(db))
		authorized.POST("/accounts", controllers.CreateAccount(db))
		authorized.PUT("/accounts/:id", controllers.UpdateAccount(db))
		authorized.DELETE("/accounts/:id", controllers.DeleteAccount(db))

		authorized.GET("/categories", controllers.GetCategories(db))
		authorized.POST("/categories", controllers.CreateCategory(db))
		authorized.PUT("/categories/:id", controllers.UpdateCategory(db))
		authorized.DELETE("/categories/:id", controllers.DeleteCategory(db))

		authorized.GET("/transactions", controllers.GetTransactions(db))
		authorized.POST("/transactions", controllers.CreateTransaction(db))
		authorized.PUT("/transactions/:id", controllers.UpdateTransaction(db))
		authorized.DELETE("/transactions/:id", controllers.DeleteTransaction(db))

		authorized.GET("/loans", controllers.GetLoans(db))
		authorized.POST("/loans", controllers.CreateLoan(db))
		authorized.PUT("/loans/:id", controllers.UpdateLoan(db))
		authorized.DELETE("/loans/:id", controllers.DeleteLoan(db))
	}
}
