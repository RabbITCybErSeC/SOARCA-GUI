package routes

import (
	"soarca-gui/handlers"
	"soarca-gui/views"

	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	publicRoutes := app.Group("/")
	PublicRoutes(publicRoutes)
}

func PublicRoutes(app *gin.RouterGroup) {
	authHandler := handlers.AuthHandler{}

	publicRoute := app.Group("/")
	{
		publicRoute.GET("/", authHandler.AuthPage)
		publicRoute.POST("/login", authHandler.Login)
		publicRoute.GET("/dashboard", handlers.HomeDashboard)
		publicRoute.GET("/reporting", handlers.ReportingDashboard)
	}
	publicRoute.StaticFS("/public", views.GetPublicAssetsFileSystem())
}
