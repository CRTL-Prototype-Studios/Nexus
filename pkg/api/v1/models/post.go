package models

type Post struct {
	BaseModel
	Content  string    `gorm:"type:text" json:"content"`
	Comments []Comment `gorm:"many2many:post_comments;" json:"comments"`
}
