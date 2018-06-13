package controllers

import (
	"SemiRevel/app/helpers"
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
	return c.Render(id, grade)
}

func (c User) Mypage() revel.Result {
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

func (c User) Input() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	return c.Render(id, grade)
}

func (c User) Password() revel.Result {
	id := c.Session["id"]
	//grade := c.Session["grade"]

	user := models.User{}
	DB.Where("id = ?", id).First(&user)

	password := helpers.ToHash(c.Params.Form.Get("password"))
	newpassword1 := c.Params.Form.Get("new_password1")
	newpassword2 := c.Params.Form.Get("new_password2")

	if password != user.Password {
		c.Flash.Error("passwordが違います")
		fmt.Println("passwordが違います")
		return c.Redirect(User.Input)

	} else if newpassword1 != newpassword2 {
		c.Flash.Error("新しいpasswordに間違いがあります．")
		fmt.Println("新しいpasswordに間違いがあります．")
		return c.Redirect(User.Input)
	} else {

		newpassword1 = helpers.ToHash(newpassword1)
		DB.Model(&user).Update("password", newpassword1).Where("id = ?", id)
		c.Flash.Success("passwordが変更できました．")
	}

	return c.Redirect(User.Index)
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
	revel.InterceptFunc(CheckUser, revel.BEFORE, &User{})
}
