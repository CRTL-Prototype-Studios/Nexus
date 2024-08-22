package main

import (
	"log"
	"nexus/pkg/api/routes"
	"nexus/pkg/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal("Error connecting to config: ", err)
	}

	minioClient, err := config.InitMinIO()
	if err != nil {
		log.Fatal("Error initializing MinIO client: ", err)
	}

	router := gin.Default()

	routes.SetupRoutes(router, db, minioClient)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
