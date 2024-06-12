package handlers

import (
	"soarca-gui/views/defaults"

	"github.com/gin-gonic/gin"
)

func Homeindex(context *gin.Context) {
	component := defaults.Layout("SOARCA", "test", "test")
	component.Render(context, context.Writer)
}
