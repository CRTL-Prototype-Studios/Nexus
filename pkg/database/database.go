package database

import (
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {

	err := NewSource()
	if err != nil {
		return nil, err
	}

	return Inst, nil
}
