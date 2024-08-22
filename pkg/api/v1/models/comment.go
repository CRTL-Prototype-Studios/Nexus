package models

type Comment struct {
	BaseModel
	Content    string `json:"content"`
	BlogPostID uint   `json:"blogPostID"`
	PostID     uint   `json:"postID"`
}
