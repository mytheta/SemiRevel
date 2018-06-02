package controllers

import (
	"SemiRevel/app/models"
	"fmt"
	"path/filepath"

	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	materials := []MaterialJoinsUser{}
	DB.Where("id = ?", id).Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("material_id desc").Limit(100).Scan(&materials)
	for n, material := range materials {
		material.File_path = filepath.Join(material.File_path, material.File_name)
		materials[n] = material

	}

	user := models.User{}
	DB.Where("id = ?", id).First(&user)
	thesis := user.Thesis

	return c.Render(materials, id, grade, thesis)
}

func (c User) Thesis() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	users := []models.User{}
	DB.Table("users").Select("users.name, users.thesis").Scan(&users)

	return c.Render(users, id, grade)
}

func (c User) UpdateThesis() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	thesis := c.Params.Form.Get("thesis")
	user := models.User{}
	DB.Model(&user).Where("id = ?", id).Update("thesis", thesis)

	return c.Render(id, grade)
}

func (c User) UpdateIndex() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]

	return c.Render(id, grade)
}

func init() {
	fmt.Println("なぜこない")
	revel.InterceptFunc(CheckUser, revel.BEFORE, &User{})
}
