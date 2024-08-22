package models

import (
	"time"
)

type BlogPost struct {
	BaseModel
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments" gorm:"-"` // Use gorm:"-" to exclude from database operations
	CoverID  *uint     `json:"coverID"`
	Cover    *Photo    `json:"cover" gorm:"foreignKey:CoverID"`
}

type Comment struct {
	BaseModel
	Content    string    `json:"content"`
	BlogPostID *uint     `json:"blogPostID"`
	BlogPost   *BlogPost `json:"blogPost" gorm:"foreignKey:BlogPostID"`
}

type Album struct {
	BaseModel
	Name   string  `json:"name"`
	Photos []Photo `json:"photos" gorm:"many2many:album_photos;"`
}

type File struct {
	BaseModel
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	Size        int64     `json:"size"`
	URL         string    `json:"url"`
	ContentType string    `json:"contentType"`
	UploadedAt  time.Time `json:"uploadedAt"`
	IsDirectory bool      `json:"isDirectory"`
}

type Photo struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	FileID      uint   `json:"fileID"`
	File        File   `json:"file" gorm:"foreignKey:FileID"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}
