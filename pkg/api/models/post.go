package models

type Post struct {
	BaseModel
	Content  string
	Comments []Comment
}
