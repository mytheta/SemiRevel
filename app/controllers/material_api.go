package controllers

import (
	"SemiRevel/app/helpers"
	"SemiRevel/app/models"
	"SemiRevel/app/routes"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	materials := []MaterialJoinsUser{}
	DB.Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("material_id desc").Limit(10).Scan(&materials)
	for n, material := range materials {
		material.File_path = filepath.Join(material.File_path, material.File_name)
		materials[n] = material
		fmt.Println(material.File_path)
	}

	return c.Render(materials, id, grade)

}

func (c MaterialApi) IndexMaterial() revel.Result {

	return c.Render(c.Session["grade"], c.Session["id"])
}

func (c MaterialApi) GradeMaterials() revel.Result {
	fmt.Println(c.Params.Route.Get("grade"))
	id := c.Session["id"]
	grade := c.Session["grade"]
	selectgrade := c.Params.Route.Get("grade")
	materials := []MaterialJoinsUser{}
	DB.Where("grade = ?", selectgrade).Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("material_id desc").Limit(100).Scan(&materials)
	for n, material := range materials {
		material.File_path = filepath.Join(material.File_path, material.File_name)
		materials[n] = material
		fmt.Println(material.File_path)
	}

	return c.Render(materials, id, grade)
}

func (c MaterialApi) SelectGrade() revel.Result {
	return c.Render(c.Session["id"])
}

func (c MaterialApi) PostMaterial(file *os.File) revel.Result {

	//時間を取得
	year, _ := strconv.Atoi(c.Params.Form.Get("year"))
	month, _ := strconv.Atoi(c.Params.Form.Get("month"))
	day, _ := strconv.Atoi(c.Params.Form.Get("day"))

	// ルーティングで設定したurlに含まれる :id とかの部分はc.Params.Route.Getで取得
	grade := c.Params.Route.Get("grade")
	id := c.Params.Route.Get("user_id")

	user := models.User{}
	DB.Where("id = ?", id).First(&user)
	userName := user.Name

	materialName := c.Params.Form.Get("material_name")
	comment := c.Params.Form.Get("comment")

	// アップロードしたファイルのファイル名を取得
	fileName := c.Params.Files["file"][0].Filename

	c.Validation.Required(materialName).Message("この項目は必須項目です").Key("material_name")
	c.Validation.Required(comment).Message("この項目は必須項目です").Key("comment")
	c.Validation.Required(fileName).Message("資料を登録してください").Key("file")
	if c.Validation.HasErrors() {
		// Store the validation errors in the flash context and redirect.
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(MaterialApi.IndexMaterial)
	}
	// 現在のディレクトリを取得
	pwd, _ := os.Getwd()

	//アップロードしたファイルの取得
	extension := strings.LastIndex(fileName, ".")

	//ファイル名をランダムの数値に変換
	randomName := helpers.Random()
	randomName += fileName[extension:]

	//fileを連結 (/Users/yutsukimiyashita/dev/src/SemiRevel/materials/grade/id/)
	gradeID := filepath.Join(grade, id)
	materialsPATH := filepath.Join("materials", gradeID)
	createPATH := filepath.Join(pwd, materialsPATH)

	err := os.MkdirAll(materialsPATH, 0777)
	fmt.Println(err)

	materialsPATH = filepath.Join("/SemiRevel", materialsPATH)

	//uploadedfileディレクトリに受け取ったファイル名でファイルを作成
	uploadedFile, err := os.Create(createPATH + "/" + randomName)
	fmt.Printf("imgFile => %v\n", uploadedFile)
	if err != nil {
		fmt.Println(err)
	}

	//作ったファイルに読み込んだファイルの内容を丸ごとコピー
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		panic(err)
	}

	// materialモデルに値を格納
	material := &models.Material{
		Material_name: materialName,
		File_name:     randomName,
		User_id:       c.Params.Route.Get("user_id"),
		Year:          year,
		Month:         month,
		Day:           day,
		File_path:     materialsPATH,
		Comment:       comment,
	}

	DB.Create(material)

	helpers.Mail(userName, materialName)

	return c.Render()
}

func (c MaterialApi) File() revel.Result {
	pwd, _ := os.Getwd()
	path := filepath.Join("/materials/", c.Params.Route.Get("grade"))
	path = filepath.Join(path, c.Params.Route.Get("user_id"))
	path = filepath.Join(path, c.Params.Route.Get("file_path"))
	path = filepath.Join(pwd, path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("エラーだよ")
	}
	fmt.Println(path)
	return c.RenderFile(f, revel.Inline)
}

func (c MaterialApi) DeleteMaterial() revel.Result {

	id := c.Session["id"]
	grade := c.Session["grade"]

	return c.Render(id, grade)
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
