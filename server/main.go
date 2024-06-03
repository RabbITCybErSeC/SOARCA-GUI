package main

import (
	defaults "soarca-gui/views/defaults"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	component := defaults.Layout("John")

	router.GET("/", func(context *gin.Context) {
		component.Render(context, context.Writer)
	})
	// // http.Handle("/", templ.Handler(component))

	// fmt.Println("Listening on :3000")
	// http.ListenAndServe(":3000", nil)
	router.Run(":8080")
}
