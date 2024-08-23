package main

import (
	"log"
	"nexus/pkg/api/v1/routes"
	"nexus/pkg/config"
	"nexus/pkg/database"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	enforceMigration := os.Getenv("ENFORCE_SCHEMA_MIGRATION")

	if err = database.InitalizeDatabase(); err != nil {
		log.Fatal("Error initializing PostgreSQL Database Instance: ", err)
	}

	if enforceMigration == "true" {
		if err := database.RunMigration(database.Inst); err != nil {
			log.Fatal("Error running database migration: ", err)
		}
		log.Println("Schema migration enforced and completed successfully.")
	} else {
		log.Println("Skipping schema migration as ENFORCE_SCHEMA_MIGRATION is not set to 'true'")
	}

	minioClient, err := config.InitializeStorage()
	if err != nil {
		log.Fatal("Error initializing MinIO client: ", err)
	}

	router := gin.Default()

	routes.SetupRoutes(router, database.Inst, minioClient)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
