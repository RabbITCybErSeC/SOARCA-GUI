package main

import (
	"soarca-gui/handlers"

	"github.com/gin-gonic/gin"
)

var (
	Version   string
	Buildtime string
)

func main() {
	app := gin.Default()
	handlers.Setup(app)
	app.Run(":8080")
}
