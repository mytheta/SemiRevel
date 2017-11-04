package models

type Material struct {
	Material_id   uint64 `gorm:"primary_key" json:"material_id"`
	Material_name string `json:"material_name"`
	User_id       string `json:"user_id"`
	Year          string `json:"year"`
	Month         uint64 `json:"month"`
	Day           uint64 `json:"day"`
	Material_type uint64 `json:"material_type"`
	File_path     string `json:"file_path"`
	comment       string `json:"comment"`
}
