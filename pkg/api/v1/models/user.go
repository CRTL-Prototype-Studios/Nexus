package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	UUID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"uuid"`
	Username string    `gorm:"uniqueIndex" json:"username"`
	Password string    `json:"-"`
	RoleID   uint      `json:"roleId"`
	Role     Role      `json:"role"`
}

type Role struct {
	BaseModel
	Name        string       `gorm:"uniqueIndex" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions"`
}

type Permission struct {
	BaseModel
	Name string `gorm:"uniqueIndex" json:"name"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UUID = uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
