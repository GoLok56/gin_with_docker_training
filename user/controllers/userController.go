package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Login Check for user's username and password with a registered user in database
func Login(data ...interface{}) (int, map[string]interface{}) {
	mapData := cast.ToStringMap(data[0])

	if mapData["username"] == "golok" && mapData["password"] == "123456" {
		return http.StatusOK, gin.H{
			"error":   "false",
			"message": "Welcome",
			"status":  http.StatusOK,
			"data":    mapData,
		}
	}

	return http.StatusUnauthorized, gin.H{
		"error":   true,
		"message": "Check username/password",
		"status":  http.StatusUnauthorized,
	}
}
