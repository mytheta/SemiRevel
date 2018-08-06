package controllers

import (
	"SemiRevel/app/daos"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Home() revel.Result {

	id := c.Session["id"]
	grade := c.Session["grade"]
	materials := daos.ShowMaterialLimitTen()

	return c.Render(materials, id, grade)

}
