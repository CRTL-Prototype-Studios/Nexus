package models

type Photo struct {
	BaseModel
	Title       string
	Description string
	ImageURL    string
	Comments    []Comment
}
