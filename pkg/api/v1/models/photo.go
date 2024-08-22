package models

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
