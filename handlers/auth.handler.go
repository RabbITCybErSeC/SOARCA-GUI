package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"soarca-gui/utils"
	authviews "soarca-gui/views/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func (a *AuthHandler) AuthPage(context *gin.Context) {
	render := utils.NewTempl(context, http.StatusOK, authviews.LoginIndex("SOARCA"))
	context.Render(http.StatusOK, render)
}

func (a *AuthHandler) Login(context *gin.Context) {
	email := context.PostForm("email")
	password := context.PostForm("password")

	var errors []string

	if email == "" {
		errors = append(errors, "Email is required")
	}

	if password == "" || len(password) < 8 {
		errors = append(errors, "Password is required")
	}

	if len(errors) > 0 {
		// template := utils.NewTempl(ctx, http.StatusOK, login_view.AuthErrorCmp(errors))
		// ctx.Render(http.StatusOK, template)
		return
	}
	context.Header("HX-Redirect", "/")
	context.String(http.StatusFound, "")
}

func (a *AuthHandler) InputValidation(context *gin.Context) {
	email := context.PostForm("email")
	password := context.PostForm("password")

	fmt.Println(email, password, "here")

	if email != "" && !strings.Contains(email, "@") {
		context.String(http.StatusOK, "Email is required")
		return
	}

	if password != "" && len(password) < 8 {
		context.String(http.StatusOK, "Password is required and must be at least 8 characters long")
		return
	}
	context.String(http.StatusOK, "")
}
