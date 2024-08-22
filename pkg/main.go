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
		log.Fatal("Error loading .env file")
	}

	enforceMigration := os.Getenv("ENFORCE_SCHEMA_MIGRATION")

	if err = database.InitalizeDatabase(); err != nil {
		log.Fatal("Error initializing PostgreSQL Database Instance: ", err)
	}

	if err := database.RunMigration(database.Inst); err != nil {
		log.Fatal("Error running database migration to PostgreSQL Database: ", err)
	} else if enforceMigration == "true" {
		log.Println("Schema migration enforced and completed successfully.")
	}

	minioClient, err := config.InitializeStorage()
	if err != nil {
		log.Fatal("Error initializing MinIO client: ", err)
	}

	router := gin.Default()

	routes.SetupRoutes(router, database.Inst, minioClient)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
