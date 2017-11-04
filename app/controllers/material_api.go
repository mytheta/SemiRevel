package controllers

import (
	"SemiRevel/app/models"

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
	response.Response = materials // 結果を格納してあげる

	return c.RenderJSON(response)
}

func (c MaterialApi) GetMaterial() revel.Result {

	response := JsonResponse{}
	response.Response = "single article"

	return c.RenderJSON(response)
}

func (c MaterialApi) PostMaterial() revel.Result {

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
