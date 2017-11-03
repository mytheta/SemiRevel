package controllers

import (
    "github.com/revel/revel"
    "SemiRevel/app/models"
)

type Authentication struct {
    *revel.Controller
}

func (c Authentication) Login() revel.Result {

    id := c.Params.Form.Get("id")
    //password := c.Params.Form.Get("password")

    user := []models.User{}
    DB.Where("id = ?", id).First(&user)

       response := JsonResponse{}
       // この時点でarticleにはidが振られているのでそのまま返してあげます
       response.Response = user

       return c.RenderJSON(response)
}
