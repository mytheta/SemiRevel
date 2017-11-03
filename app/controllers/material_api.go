package controllers

import (
    "github.com/revel/revel"
    "SemiRevel/app/models"
)

type MaterialApi struct {
    *revel.Controller
}

func (c MaterialApi) GetMaterials() revel.Result {

    materials := []models.Material{}
    DB.Order("id sec").Find(&materials)

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
