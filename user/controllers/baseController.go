package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../models"
)

// BaseController A function that will return a handler for routes, also sending response to the client
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
