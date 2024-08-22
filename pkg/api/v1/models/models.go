package models

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
	PostID     *uint     `json:"postID"`
	Post       *Post     `json:"post" gorm:"foreignKey:PostID"`
}

type Photo struct {
	BaseModel
	Title        string `json:"title"`
	Description  string `json:"description"`
	Aperture     int8   `json:"aperture"`
	ShutterSpeed int8   `json:"shutter_speed"`
	ISO          int16  `json:"iso"`
	Lens         string `json:"lens"`
	ImageURL     string `json:"image_url"`
}

type Post struct {
	BaseModel
	Content  string    `json:"content"`
	Comments []Comment `json:"comments"`
}

type Album struct {
	BaseModel
	Name   string  `json:"name"`
	Photos []Photo `json:"photos" gorm:"many2many:album_photos;"`
}
