package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize config connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Error connecting to config: ", err)
	}

	// Initialize MinIO client
	minioClient, err := config.InitMinIO()
	if err != nil {
		log.Fatal("Error initializing MinIO client: ", err)
	}

	// Create a new Gin router
	router := gin.Default()

	// Initialize routes
	routes.SetupRoutes(router, db, minioClient)

	// Start the server
	router.Run(":8080")
}
