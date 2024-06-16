package handlers

import (
	"net/http"

	"soarca-gui/utils"
	"soarca-gui/views"
	"soarca-gui/views/dashboard/reporting"

	"github.com/gin-gonic/gin"
)

func HomeDashboard(context *gin.Context) {
	render := utils.NewTempl(context, http.StatusOK, views.Home(nil))
	context.Render(http.StatusOK, render)
}

func ReportingDashboard(context *gin.Context) {

	render := utils.NewTempl(context, http.StatusOK,
		reporting.ReportingIndex())

	context.Render(http.StatusOK, render)
}

func ReportingCard(context *gin.Context) {

	render := utils.NewTempl(context, http.StatusOK, reporting.ReportingIndex())

	context.Render(http.StatusOK, render)
}
