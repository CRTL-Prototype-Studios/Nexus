package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Inst *gorm.DB

func InitalizeDatabase() error {
	var err error

	dbHost, dbUser, dbPassword, dbName, dbPort :=
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	Inst, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return err
}
