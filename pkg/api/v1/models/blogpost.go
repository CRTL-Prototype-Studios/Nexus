package models

type BlogPost struct {
	BaseModel
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
	Cover    Photo     `json:"cover"`
}
