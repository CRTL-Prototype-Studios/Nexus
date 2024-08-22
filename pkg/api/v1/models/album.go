package models

type Album struct {
	BaseModel
	Name   string  `json:"name"`
	Photos []Photo `json:"photos"`
}
