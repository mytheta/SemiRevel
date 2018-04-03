package models

type User struct {
	Id       string `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Grade    string `json:"grade"`
	Role     int    `json:"role"`
}
