package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(data map[string]interface{}) (int, map[string]interface{}) {
	if data["username"] == "golok" && data["password"] == "123456" {
		return http.StatusOK, gin.H{
			"error":   "false",
			"message": "Welcome",
			"status":  http.StatusOK,
			"data":    gin.H{"username": data["username"], "password": data["password"]},
		}
	} else {
		return http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Check username/password",
			"status":  http.StatusUnauthorized,
		}
	}
}
