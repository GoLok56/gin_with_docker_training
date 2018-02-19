package repositories

import (
	"github.com/golok56/gin_with_docker_training/user/config"
	"github.com/golok56/gin_with_docker_training/user/models"
)

func Create(user models.User) models.User {
	db := config.OpenConnection()
	defer db.Close()

	db.Create(&user)

	return user
}

func FindUserByUsername(username string) models.User {
	db := config.OpenConnection()
	defer db.Close()

	user := models.User{}
	db.Where("username = ?", username).First(&user)

	return user
}

func UpdateUser(user *models.User, data map[string]interface{}) {
	db := config.OpenConnection()
	defer db.Close()

	db.Model(&user).Update(data)
}

func DeleteUser(user models.User) {
	db := config.OpenConnection()
	defer db.Close()

	db.Delete(&user)
}
