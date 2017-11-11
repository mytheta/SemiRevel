package controllers

import (
	"SemiRevel/app/models"
	"fmt"
	"os"

	"github.com/revel/revel"
)

type MaterialApi struct {
	*revel.Controller
}

type MaterialJoinsUser struct {
	models.Material
	models.User
}

func (c MaterialApi) GetMaterials() revel.Result {

	materials := []MaterialJoinsUser{}
	DB.Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("id desc").Limit(100).Scan(&materials)

	response := JsonResponse{}
	response.Response = materials

	return c.RenderJSON(response)
}

func (c MaterialApi) GetMaterial() revel.Result {

	grade := c.Params.Route.Get("grade")

	materials := []MaterialJoinsUser{}
	DB.Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Where("users.grade = ?", grade).Order("id desc").Limit(100).Scan(&materials)

	response := JsonResponse{}
	response.Response = materials

	return c.RenderJSON(response)
}

func (c MaterialApi) PostMaterial() revel.Result {

	pwd, _ := os.Getwd()

	err := os.Mkdir("materials", 0777)
	fmt.Println(err)

	fmt.Println(pwd)

	response := JsonResponse{}
	response.Response = "post article"

	return c.RenderJSON(response)
}

func (c MaterialApi) ViewMaterial() revel.Result {

	response := JsonResponse{}
	response.Response = "put article"

	return c.RenderJSON(response)
}

func (c MaterialApi) DeleteMaterial() revel.Result {

	response := JsonResponse{}
	response.Response = "delete article"

	return c.RenderJSON(response)
}
