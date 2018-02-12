package routers

import (
	"github.com/gin-gonic/gin"

	"../controllers"
)

func InitRoute(app *gin.Engine) {
	user := app.Group("/user")
	{
		user.POST("/login", controllers.BaseController(controllers.Login))
	}
}
