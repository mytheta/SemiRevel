package controllers

import (
	"SemiRevel/app/models"
	"fmt"
	"io"
	"myapp/app/routes"
	"os"
	"path/filepath"
	"strings"
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

	id := c.Session["id"]
	grade := c.Session["grade"]

	materials := []MaterialJoinsUser{}
	DB.Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("material_id desc").Limit(100).Scan(&materials)
	for n, material := range materials {
		material.File_path = filepath.Join(material.File_path, material.File_name)
		materials[n] = material
		fmt.Println(material.File_path)
	}

	for _, material := range materials {
		fmt.Println(material.File_path)
	}

	return c.Render(materials, id, grade)
}

func (c MaterialApi) IndexMaterial() revel.Result {
	id := c.Session["id"]
	grade := c.Session["grade"]
	fmt.Println("aaaaaaaa")
	fmt.Println(id)

	return c.Render(grade, id)
}

func (c MaterialApi) GradeMaterials() revel.Result {

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

	for _, material := range materials {
		fmt.Println(material.File_path)
	}

	return c.Render(id, grade)
}

func (c MaterialApi) MyMaterials() revel.Result {

	id := c.Session["id"]
	grade := c.Session["grade"]
	materials := []MaterialJoinsUser{}
	DB.Where("id = ?", id).Table("materials").Select("materials.*, users.name, users.id").Joins("INNER JOIN users ON users.id = materials.user_id").Order("material_id desc").Limit(100).Scan(&materials)
	for n, material := range materials {
		material.File_path = filepath.Join(material.File_path, material.File_name)
		materials[n] = material
		fmt.Println(material.File_path)
	}

	for _, material := range materials {
		fmt.Println(material.File_path)
	}

	return c.Render(materials, id, grade)
}

func (c MaterialApi) SelectGrade() revel.Result {

	return c.Render()
}

func (c MaterialApi) PostMaterial(file *os.File) revel.Result {

	//時間を取得
	time := time.Now()
	year, month, date := time.Date()
	intMonth := int(month)

	// ルーティングで設定したurlに含まれる :id とかの部分はc.Params.Route.Getで取得
	grade := c.Params.Route.Get("grade")
	id := c.Params.Route.Get("user_id")

	materialName := c.Params.Form.Get("material_name")
	comment := c.Params.Form.Get("comment")

	// 現在のディレクトリを取得
	pwd, _ := os.Getwd()

	//materialテーブルの最後のautoincrementを取ってくる．
	materialID := models.Material{}
	DB.Last(&materialID)

	// アップロードしたファイルのファイル名を取得
	fileName := c.Params.Files["file"][0].Filename

	//アップロードしたファイルの取得
	extension := strings.LastIndex(fileName, ".")

	//ファイル名をランダムの数値に変換
	// var randomName int
	// randomName, _ = strconv.Atoi(random())
	randomName := random()
	randomName += fileName[extension:]

	//stringに変換
	// randomStringName := strconv.Itoa(randomName)

	fmt.Println(materialID.Material_id)

	//fileを連結 (/Users/yutsukimiyashita/dev/src/SemiRevel/materials/grade/id/)
	gradeID := filepath.Join(grade, id)
	materialsPATH := filepath.Join("materials", gradeID)
	createPATH := filepath.Join(pwd, materialsPATH)

	err := os.MkdirAll(materialsPATH, 0777)
	fmt.Println(err)

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
		Month:         intMonth,
		Day:           date,
		File_path:     materialsPATH,
		Comment:       comment,
	}

	DB.Create(material)

	response := JsonResponse{}
	response.Response = material

	//メール機能
	// Connect to the remote SMTP server.
	// d, err := smtp.Dial("sapphire.u-gakugei.ac.jp:25")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Set the sender and recipient.
	// d.Mail("SemiRevel@sapphire.u-gakugei.ac.jp") // メールの送り主を指定
	// d.Rcpt("hazelab@sapphire.u-gakugei.ac.jp")   // 受信者を指定
	//
	// // Send the email body.
	// wc, err := d.Data()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer wc.Close()
	// //ToにするかCcにするかBccにするかはDATAメッセージ次第
	// buf := bytes.NewBufferString("To:hazelab@sapphire.u-gakugei.ac.jp")
	// buf.WriteString("\r\n") // DATA メッセージはCRLFのみ
	// buf.WriteString("\r\n")
	// buf.WriteString("ゼミ資料管理システム") //件名
	// buf.WriteString("\r\n")
	// buf.WriteString("新しい資料が登録されました")
	// if _, err = buf.WriteTo(wc); err != nil {
	// 	log.Fatal(err)
	// }

	// d.Quit() //メールセッションの終了

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

func checkUser(c *revel.Controller) revel.Result {
	fmt.Println("checkuser")
	if id, ok := c.Session["id"]; ok != true {
		fmt.Println(id)
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Index())
	}

	return nil
}

func init() {
	// revel.InterceptFunc(checkUser, revel.BEFORE, &MaterialApi{})
}
