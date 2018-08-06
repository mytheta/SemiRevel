package daos

import (
	"SemiRevel/app/models"
	"fmt"
	"os"
	"path/filepath"
)

type MaterialJoinsUser struct {
	models.Material
	models.User
}

func Create(material models.Material) {
	fmt.Println(material.Material_name)
	DB.Create(&material)
}

func ShowMaterialsByGrade(grade string) []MaterialJoinsUser {
	var materials []MaterialJoinsUser
	DB.Where("grade = ?", grade).Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("material_id desc").Limit(100).Scan(&materials)

	for n, material := range materials {
		material.File_path = filepath.Join(material.File_path, material.File_name)
		materials[n] = material
		fmt.Println(material.File_path)
	}

	return materials
}

func ShowMaterial(material models.Material) models.Material {
	DB.First(&material)
	return material
}

func ShowMaterialLimitTen() []MaterialJoinsUser {
	materials := []MaterialJoinsUser{}
	DB.Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("material_id desc").Limit(10).Scan(&materials)
	fmt.Println(materials)
	for n, material := range materials {
		fmt.Println("select materials")
		material.File_path = filepath.Join(material.File_path, material.File_name)
		materials[n] = material
		fmt.Println(material.File_path)
	}

	return materials
}

func Edit(material models.Material, materialName, comment string) {
	material = ShowMaterial(material)
	material.Material_name = materialName
	material.Comment = comment
	DB.Save(&material)

}
func Delete(material models.Material) {
	DB.First(&material)

	pwd, _ := os.Getwd()
	createPATH := filepath.Join(pwd, material.File_path)
	if err := os.Remove(createPATH + "/" + material.File_name); err != nil {
		fmt.Println(err)
	}

	DB.Delete(&material)
}

func MyMaterials(id string) []MaterialJoinsUser {
	materials := []MaterialJoinsUser{}
	DB.Where("id = ?", id).Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("material_id desc").Limit(100).Scan(&materials)
	for n, material := range materials {
		material.File_path = filepath.Join(material.File_path, material.File_name)
		materials[n] = material

	}

	return materials

}
