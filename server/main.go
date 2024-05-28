package main

import (
	"fmt"
	"net/http"

	defaults "soarca-gui/views/defaults"

	"github.com/a-h/templ"
)

func main() {
	component := defaults.Layout("John")
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
