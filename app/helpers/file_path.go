package helpers

import (
	"os"
	"path/filepath"
)

func GetPath(user_id, filePATH string) string {
	pwd, _ := os.Getwd()
	path := filepath.Join("/materials/", user_id)
	path = filepath.Join(path, filePATH)
	path = filepath.Join(pwd, path)

	return path
}
