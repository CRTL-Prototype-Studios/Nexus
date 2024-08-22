package database

import (
	"gorm.io/gorm"
	models2 "nexus/pkg/api/v1/models"
)

var AutoMaintainRange = []any{
	&models2.Post{},
	&models2.Photo{},
	&models2.Comment{},
	&models2.BlogPost{},
}

func RunMigration(source *gorm.DB) error {
	if err := source.AutoMigrate(
		AutoMaintainRange...,
	); err != nil {
		return err
	}

	return nil
}
