package controllers

import (
	"SemiRevel/app/models"
	"SemiRevel/app/routes"
	"fmt"

	"github.com/revel/revel"
)

type Authentication struct {
	*revel.Controller
}

func (c Authentication) Login() revel.Result {
	id := c.Params.Form.Get("id")
	password := toHash(c.Params.Form.Get("password"))
	user := models.User{}
	DB.Where("id = ?", id).First(&user)

	if password == user.Password {
		c.Session["id"] = id
		c.Session["grade"] = user.Grade
		fmt.Println("認証成功")
	} else {
		c.Flash.Error("パスワードが違います．")
		return c.Redirect(routes.App.Index())
	}
	return c.Redirect(routes.MaterialApi.Home())
}

func (c Authentication) Logout() revel.Result {

	delete(c.Session, "id")    // Removed item from session
	delete(c.Session, "grade") // Removed item from session
	c.Flash.Success("logoutしました")
	return c.Redirect(routes.App.Index())
}
