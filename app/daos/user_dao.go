package daos

import (
	"SemiRevel/app/models"
)

func ShowUserName(id string) string {
	var user models.User
	DB.Where("id = ?", id).First(&user)
	return user.Name
}

func ShowUser(id string) models.User {
	user := models.User{}
	DB.Where("id = ?", id).First(&user)
	return user

}

func UpdatePassword(id, password string) {
	user := models.User{}
	DB.Model(&user).Update("password", password).Where("id = ?", id)
}

func ShowThesis() []models.User {
	users := []models.User{}
	DB.Table("users").Select("users.name, users.thesis").Scan(&users)
	return users
}

func UpdateThesis(id, thesis string) {
	user := models.User{}
	DB.Model(&user).Where("id = ?", id).Update("thesis", thesis)

}
