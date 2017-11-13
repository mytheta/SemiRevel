package models

type Material struct {
	Material_id   int    `gorm:"primary_key" json:"material_id"`
	Material_name string `json:"material_name"`
	User_id       string `json:"user_id"`
	Year          int    `json:"year"`
	Month         int    `json:"month"`
	Day           int    `json:"day"`
	Material_type string `json:"material_type"`
	File_path     string `json:"file_path"`
	Comment       string `json:"comment"`
}
