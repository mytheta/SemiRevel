package controllers

import (
	"SemiRevel/app/daos"
	"SemiRevel/app/helpers"
	"SemiRevel/app/models"
	"SemiRevel/app/routes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/revel/revel"
)

type MaterialApi struct {
	*revel.Controller
}

type MaterialJoinsUser struct {
	models.Material
	models.User
}

func (c MaterialApi) Home() revel.Result {

	id := c.Session["id"]
	grade := c.Session["grade"]
	materials := daos.ShowMaterialLimitTen()

	return c.Render(materials, id, grade)

}

func (c MaterialApi) Index() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	return c.Render(id, grade)
}

func (c MaterialApi) Create(file *os.File) revel.Result {

	//時間を取得
	year, month, day := helpers.ConvertStringToInt(c.Params.Form.Get("year"), c.Params.Form.Get("month"), c.Params.Form.Get("day"))
	materialName := c.Params.Form.Get("material_name")
	comment := c.Params.Form.Get("comment")
	fileName := c.Params.Files["file"][0].Filename

	grade := c.Session["grade"]
	id := c.Session["id"]

	//user := daos.ShowUser(id)
	userName := daos.ShowUserName(id)

	c.Validation.Required(materialName).Message("この項目は必須項目です").Key("material_name")
	c.Validation.Required(comment).Message("この項目は必須項目です").Key("comment")
	c.Validation.Required(fileName).Message("資料を登録してください").Key("file")
	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(MaterialApi.Index)
	}

	filePath := helpers.MkdirMaterialPath(id)

	extension := strings.LastIndex(fileName, ".")
	randomName := helpers.Random()
	randomName += fileName[extension:]

	helpers.MakeFile(filePath, randomName, file)

	// materialモデルに値を格納
	material := models.Material{
		Material_name: materialName,
		File_name:     randomName,
		User_id:       id,
		Year:          year,
		Month:         month,
		Day:           day,
		File_path:     filePath,
		Comment:       comment,
	}

	daos.Create(material)

	helpers.Mail(userName, materialName)

	return c.Render(id, grade)
}

func (c MaterialApi) GradeMaterials() revel.Result {
	fmt.Println(c.Params.Route.Get("grade"))
	id := c.Session["id"]
	grade := c.Session["grade"]
	selectgrade := c.Params.Route.Get("grade")
	materials := daos.ShowMaterialsByGrade(selectgrade)

	return c.Render(materials, id, grade)
}

func (c MaterialApi) Delete() revel.Result {
	// id := c.Session["id"]
	// grade := c.Session["grade"]
	material := models.Material{}
	material.Material_id, _ = strconv.Atoi(c.Params.Route.Get("id"))
	daos.Delete(material)

	c.Flash.Success("削除完了しました")
	return c.Redirect(routes.User.Mypage())
}

func (c MaterialApi) EditIndex() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	material_id, _ := strconv.Atoi(c.Params.Route.Get("id"))

	material := models.Material{}
	material.Material_id = material_id
	material = daos.ShowMaterial(material)
	material_name := material.Material_name
	year := material.Year
	month := material.Month
	day := material.Day
	comment := material.Comment
	fmt.Println(material_name)

	return c.Render(material_name, year, month, day, comment, id, grade, material_id)
}

func (c MaterialApi) Edit() revel.Result {

	materialName := c.Params.Form.Get("material_name")
	comment := c.Params.Form.Get("comment")

	material := models.Material{}
	material.Material_id, _ = strconv.Atoi(c.Params.Route.Get("id"))
	daos.Edit(material, materialName, comment)

	c.Flash.Success("編集完了しました")
	return c.Redirect(routes.User.Mypage())
}

func (c MaterialApi) File() revel.Result {
	path := helpers.GetPath(c.Params.Route.Get("user_id"), c.Params.Route.Get("file_path"))
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("エラーだよ")
	}
	fmt.Println(path)
	return c.RenderFile(f, revel.Inline)
}

func CheckUser(c *revel.Controller) revel.Result {
	fmt.Println("checkuser")
	if id, ok := c.Session["id"]; ok != true {
		fmt.Println(id)
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Index())
	}

	return nil
}

func init() {
	revel.InterceptFunc(CheckUser, revel.BEFORE, &MaterialApi{})
}
