package models

type Comment struct {
	BaseModel
	Content    string
	BlogPostID uint
	PhotoID    uint
	PostID     uint
}
