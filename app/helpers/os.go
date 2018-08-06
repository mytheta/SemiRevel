package helpers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func MkdirMaterialPath(id string) string {
	path := filepath.Join("materials", id)
	err := os.MkdirAll(path, 0777)
	fmt.Println(err)

	return path
}

func MakeFile(path, fileName string, file *os.File) {
	// 現在のディレクトリを取得
	pwd, _ := os.Getwd()
	path = filepath.Join(pwd, path)

	uploadedFile, err := os.Create(path + "/" + fileName)
	fmt.Printf("imgFile => %v\n", uploadedFile)
	if err != nil {
		fmt.Println(err)
	}

	//作ったファイルに読み込んだファイルの内容を丸ごとコピー
	_, err = io.Copy(uploadedFile, file)
	if err != nil {
		panic(err)
	}

}
