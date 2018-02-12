package main

import (
	"github.com/gin-gonic/gin"

	"./routers"
)

const port = ":8001"

func main() {
	app := gin.Default()

	routers.InitRoute(app)

	app.Run(port)
}
