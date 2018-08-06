package controllers

import (
	"SemiRevel/app/daos"
	"SemiRevel/app/helpers"
	"fmt"

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
	materials := daos.MyMaterials(id)

	user := daos.ShowUser(id)
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

	user := daos.ShowUser(id)

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
		daos.UpdatePassword(id, newpassword1)
		c.Flash.Success("passwordが変更できました．")
	}

	return c.Redirect(User.Index)
}

func (c User) Thesis() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	users := daos.ShowThesis()
	return c.Render(users, id, grade)
}

func (c User) UpdateThesis() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	thesis := c.Params.Form.Get("thesis")
	daos.UpdateThesis(id, thesis)

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
