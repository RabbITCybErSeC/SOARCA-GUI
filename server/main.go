package main

import (
	"soarca-gui/views"
	defaults "soarca-gui/views/defaults"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	component := defaults.Layout("John")

	router.GET("/", func(context *gin.Context) {
		component.Render(context, context.Writer)
	})

	router.StaticFS("/public", views.GetPublicAssetsFileSystem())
	router.Run(":8080")
}
