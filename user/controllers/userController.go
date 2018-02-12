package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"../config"
	"../models"
)

// Login Check for user's username and password with a registered user in database
func Login(data ...interface{}) (int, map[string]interface{}) {
	mapData := cast.ToStringMap(data[0])
	db := config.OpenConnection()
	defer db.Close()

	user := models.User{}

	db.Where("username = ?", mapData["username"]).First(&user)
	if user == (models.User{}) {
		return http.StatusUnauthorized, gin.H{
			"error":   "true",
			"status":  http.StatusUnauthorized,
			"message": "User not found!",
		}
	}

	if user.Password != mapData["password"] {
		return http.StatusUnauthorized, gin.H{
			"error":   "true",
			"status":  http.StatusUnauthorized,
			"message": "Password doesn't match!",
		}
	}

	return http.StatusOK, gin.H{
		"error":   "false",
		"status":  http.StatusOK,
		"message": "Login success, welcome " + user.Name,
	}
}

// Register Register a new user to database
func Register(data ...interface{}) (int, map[string]interface{}) {
	mapData := cast.ToStringMap(data[0])
	db := config.OpenConnection()
	defer db.Close()

	user := models.User{
		Username: mapData["username"].(string),
		Password: mapData["password"].(string),
		Name:     mapData["name"].(string),
	}

	db.Create(&user)

	return http.StatusOK, gin.H{
		"error":   false,
		"status":  http.StatusOK,
		"message": "Success",
		"data":    user,
	}
}
