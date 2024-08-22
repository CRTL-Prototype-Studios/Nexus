package models

type BlogPost struct {
	BaseModel
	Title    string
	Content  string
	Comments []Comment
	Cover    Photo
}
