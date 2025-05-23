// @title Finance Tracker API
// @version 1.0
// @description API dokumentasi untuk aplikasi keuangan
// @contact.name Tim Dev
// @contact.email kamu@email.com
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"finance-tracker/config"
	"finance-tracker/routes"
	"log"
	"os"

	_ "finance-tracker/cmd/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files" // <-- Tambah ini
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("Tidak bisa memuat .env, menggunakan environment bawaan.")
	}

	// Setup DB
	config.InitDB()
	db := config.DB
	defer db.Close()

	// Setup Router
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(router, db)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server berjalan di port %s...", port)
	router.Run(":" + port)
}
