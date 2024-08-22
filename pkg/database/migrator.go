package database

import (
	"gorm.io/gorm"
	"nexus/pkg/api/models"
)

var AutoMaintainRange = []any{
	&models.Post{},
	&models.Photo{},
	&models.Comment{},
	&models.BlogPost{},
}

func RunMigration(source *gorm.DB) error {
	if err := source.AutoMigrate(
		AutoMaintainRange...,
	); err != nil {
		return err
	}

	return nil
}
