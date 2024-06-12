package handlers

import (
	"net/http"

	"soarca-gui/utils"
	"soarca-gui/views"

	"github.com/gin-gonic/gin"
)

func Homeindex(context *gin.Context) {
	render := utils.NewTempl(context, http.StatusOK, views.Home())
	context.Render(http.StatusOK, render)
}
