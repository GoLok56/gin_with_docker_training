package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/golok56/gin_with_docker_training/user/models"
)

// BaseController A function that will return a handler for routes, also sending response to the client
func BaseController(controller models.Controller, paramName ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			statusCode    int
			requestMethod string
			body, data    map[string]interface{}
		)
		requestMethod = c.Request.Method

		err := c.ShouldBindJSON(&body)

		switch requestMethod {
		case "GET", "DELETE":
			statusCode, data = controller(c.Param(paramName[0]))
		case "POST":
			if err == nil {
				statusCode, data = controller(body)
			}
		case "PUT":
			if err == nil {
				statusCode, data = controller(c.Param(paramName[0]), body)
			}
		default:
			c.JSON(http.StatusBadRequest, err.Error)
		}

		c.JSON(statusCode, data)
	}
}
