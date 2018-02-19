package main

import (
	"github.com/gin-gonic/gin"

	"github.com/golok56/gin_with_docker_training/user/routers"
)

const port = ":8001"

func main() {
	app := gin.Default()

	routers.InitRoute(app)

	app.Run(port)
}
