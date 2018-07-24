package controllers

import (
	"fmt"
	"path/filepath"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Home() revel.Result {

	id := c.Session["id"]
	grade := c.Session["grade"]
	materials := []MaterialJoinsUser{}
	DB.Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("material_id desc").Limit(10).Scan(&materials)
	fmt.Println(materials)
	for n, material := range materials {
		fmt.Println("select materials")
		material.File_path = filepath.Join(material.File_path, material.File_name)
		materials[n] = material
		fmt.Println(material.File_path)
	}

	return c.Render(materials, id, grade)

}
