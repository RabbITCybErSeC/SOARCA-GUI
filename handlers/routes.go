package handlers

import (
	"soarca-gui/views"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	app.GET("/", Homeindex)

	app.StaticFS("/public", views.GetPublicAssetsFileSystem())
}
