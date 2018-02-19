package controllers

import (
	"net/http"

	"github.com/spf13/cast"

	"github.com/golok56/gin_with_docker_training/user/config"
	"github.com/golok56/gin_with_docker_training/user/models"
	repo "github.com/golok56/gin_with_docker_training/user/repositories"
)

// Login Check for user's username and password with a registered user in database
func Login(data ...interface{}) (int, map[string]interface{}) {
	mapData := cast.ToStringMap(data[0])
	db := config.OpenConnection()
	defer db.Close()

	user := models.User{}

	db.Where("username = ?", mapData["username"]).First(&user)
	if user == (models.User{}) {
		return http.StatusUnauthorized, map[string]interface{}{
			"error":   "true",
			"status":  http.StatusUnauthorized,
			"message": "User not found!",
		}
	}

	if user.Password != mapData["password"] {
		return http.StatusUnauthorized, map[string]interface{}{
			"error":   "true",
			"status":  http.StatusUnauthorized,
			"message": "Password doesn't match!",
		}
	}

	user.Password = ""
	return http.StatusOK, map[string]interface{}{
		"error":   "false",
		"status":  http.StatusOK,
		"message": "Login success, welcome " + user.Name,
	}
}

// Register Register a new user to database
func Register(data ...interface{}) (int, map[string]interface{}) {
	mapData := cast.ToStringMap(data[0])

	user := repo.Create(models.User{
		Username: mapData["username"].(string),
		Password: mapData["password"].(string),
		Name:     mapData["name"].(string),
	})

	return http.StatusOK, map[string]interface{}{
		"error":   false,
		"status":  http.StatusOK,
		"message": "Success",
		"data":    user,
	}
}

func Get(data ...interface{}) (int, map[string]interface{}) {
	username := data[0].(string)
	user := repo.FindUserByUsername(username)

	if user == (models.User{}) {
		return http.StatusNotFound, map[string]interface{}{
			"error":   "true",
			"status":  http.StatusNotFound,
			"message": "User not found!",
		}
	}

	user.Password = ""
	return http.StatusOK, map[string]interface{}{
		"error":   false,
		"status":  http.StatusOK,
		"message": "Success",
		"data":    user,
	}
}

func Update(data ...interface{}) (int, map[string]interface{}) {
	username := data[0].(string)
	body := cast.ToStringMap(data[1])

	user := repo.FindUserByUsername(username)
	if user == (models.User{}) {
		return http.StatusNotFound, map[string]interface{}{
			"error":   "true",
			"status":  http.StatusNotFound,
			"message": "User not found!",
		}
	}

	repo.UpdateUser(&user, body)
	user.Password = ""
	return http.StatusOK, map[string]interface{}{
		"error":   false,
		"status":  http.StatusOK,
		"message": "Success",
		"data":    user,
	}
}

func Delete(data ...interface{}) (int, map[string]interface{}) {
	username := data[0].(string)

	user := repo.FindUserByUsername(username)
	if user == (models.User{}) {
		return http.StatusNotFound, map[string]interface{}{
			"error":   "true",
			"status":  http.StatusNotFound,
			"message": "User not found!",
		}
	}

	repo.DeleteUser(user)
	user.Password = ""
	return http.StatusOK, map[string]interface{}{
		"error":   false,
		"status":  http.StatusOK,
		"message": "Success",
		"data":    user,
	}
}
