package controllers

import (
	"SemiRevel/app/models"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

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

	return c.Render(response)
}

func (c MaterialApi) GetMaterial() revel.Result {

	grade := c.Params.Route.Get("grade")

	materials := []MaterialJoinsUser{}
	DB.Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Where("users.grade = ?", grade).Order("id desc").Limit(100).Scan(&materials)

	response := JsonResponse{}
	response.Response = materials

	return c.RenderJSON(response)
}

func (c MaterialApi) PostMaterial(file *os.File) revel.Result {

	//時間を取得
	time := time.Now()
	year, month, date := time.Date()
	intMonth := int(month)

	// ルーティングで設定したurlに含まれる :id とかの部分はc.Params.Route.Getで取得
	grade := c.Params.Route.Get("grade")
	id := c.Params.Route.Get("user_id")

	// 現在のディレクトリを取得
	pwd, _ := os.Getwd()

	//materialテーブルの最後のautoincrementを取ってくる．
	materialID := models.Material{}
	DB.Last(&materialID)

	// アップロードしたファイルのファイル名を取得
	fileName := c.Params.Files["file"][0].Filename

	fmt.Println(materialID.Material_id)

	//fileを連結 (/Users/yutsukimiyashita/dev/src/SemiRevel/materials/grade/id/)
	gradeID := filepath.Join(grade, id)
	materialsPATH := filepath.Join("materials", gradeID)
	createPATH := filepath.Join(pwd, materialsPATH)

	err := os.MkdirAll(materialsPATH, 0777)
	fmt.Println(err)

	//uploadedfileディレクトリに受け取ったファイル名でファイルを作成
	uploadedFile, err := os.Create(createPATH + "/" + fileName)
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
		Material_name: fileName,
		User_id:       c.Params.Route.Get("user_id"),
		Year:          year,
		Month:         intMonth,
		Day:           date,
		Material_type: c.Params.Form.Get("material_type"),
		File_path:     materialsPATH,
		Comment:       c.Params.Form.Get("comment"),
	}

	DB.Create(material)

	response := JsonResponse{}
	response.Response = material

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
