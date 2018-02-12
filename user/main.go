package main

import (
	"github.com/gin-gonic/gin"

	"./routers"
)

const PORT = ":8001"

func main() {
	app := gin.Default()

	routers.InitRoute(app)

	app.Run(PORT)
}
