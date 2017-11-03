package controllers

import (
    "github.com/revel/revel"
    "SemiRevel/app/models"
    "fmt"
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

       return c.RenderJSON(response)
}
