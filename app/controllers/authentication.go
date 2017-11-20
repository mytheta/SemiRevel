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
	password := c.Params.Form.Get("password")

	user := models.User{}
	DB.Where("id = ?", id).First(&user)

	if password == user.Password {
		fmt.Println("認証成功")
	}

	response := JsonResponse{}
	response.Response = user

	return c.Redirect(routes.MaterialApi.GetMaterials())
}
