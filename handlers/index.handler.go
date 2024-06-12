package handlers

import (
	"net/http"

	"soarca-gui/utils"
	authviews "soarca-gui/views/auth"

	"github.com/gin-gonic/gin"
)

func Homeindex(context *gin.Context) {
	render := utils.NewTempl(context, http.StatusOK, authviews.LoginIndex("SOARCA"))
	context.Render(http.StatusOK, render)
}
