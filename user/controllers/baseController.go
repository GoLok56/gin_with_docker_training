package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../models"
)

func BaseController(controller models.Controller) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body map[string]interface{}
		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error)
			return
		}

		statusCode, data := controller(body)
		c.JSON(statusCode, data)
	}
}
