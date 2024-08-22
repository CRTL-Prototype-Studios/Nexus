package models

import (
	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model
	Title    string
	Content  string
	Comments []Comment
	Cover    Photo
}

type Photo struct {
	gorm.Model
	Title       string
	Description string
	ImageURL    string
	Comments    []Comment
}

type Post struct {
	gorm.Model
	Content  string
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content    string
	BlogPostID uint
	PhotoID    uint
	PostID     uint
}
