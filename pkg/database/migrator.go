package database

import (
	"gorm.io/gorm"
	"nexus/pkg/api/v1/models"
)

var AutoMaintainRange = []any{
	&models.Photo{},
	&models.Comment{},
	&models.BlogPost{},
	&models.Album{},
	&models.Role{},
	&models.File{},
	&models.Permission{},
	&models.User{},
}

func RunMigration(source *gorm.DB) error {
	// Enable UUID extension
	if err := source.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return err
	}

	if err := source.AutoMigrate(
		AutoMaintainRange...,
	); err != nil {
		return err
	}

	return nil
}
