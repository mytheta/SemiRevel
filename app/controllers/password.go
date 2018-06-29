package controllers

import (
	"SemiRevel/app/helpers"
	"SemiRevel/app/models"
	"fmt"

	"github.com/revel/revel"
)

type Password struct {
	*revel.Controller
}

func (c Password) Input() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	return c.Render(id, grade)
}

func (c Password) Password() revel.Result {
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

func init() {
	revel.InterceptFunc(CheckUser, revel.BEFORE, &Password{})
}
