package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/golok56/gin_with_docker_training/user/controllers"
)

func InitRoute(app *gin.Engine) {
	user := app.Group("/user")
	{
		user.POST("/login", controllers.BaseController(controllers.Login))
		user.POST("/register", controllers.BaseController(controllers.Register))
		user.GET("/:username", controllers.BaseController(controllers.Get, "username"))
		user.PUT("/:username", controllers.BaseController(controllers.Update, "username"))
		user.DELETE("/:username", controllers.BaseController(controllers.Delete, "username"))
	}
}
