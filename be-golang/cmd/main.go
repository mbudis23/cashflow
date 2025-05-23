package main

import (
	"finance-tracker/config"
	"finance-tracker/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	routes.RegisterRoutes(router, db)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server berjalan di port %s...", port)
	router.Run(":" + port)
}
