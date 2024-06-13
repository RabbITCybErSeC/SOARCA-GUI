package handlers

import (
	"errors"
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
	errors := a.inputValidation(context)
	email := context.PostForm("email")
	password := context.PostForm("password")

	fmt.Println(email, " ", password, len(errors))
	if len(errors) > 0 {
		template := utils.NewTempl(context, http.StatusOK, authviews.AuthErrorCmp(errors))
		context.Render(http.StatusOK, template)
		return
	}
	context.Header("HX-Redirect", "/")
	context.String(http.StatusFound, "")
}

func (a *AuthHandler) inputValidation(context *gin.Context) []error {
	email := context.PostForm("email")
	password := context.PostForm("password")

	var validationErrors []error
	if email == "" || !strings.Contains(email, "@") {
		validationErrors = append(validationErrors, errors.New("e-mail is required"))
	}
	if password != "" || len(password) < 8 {
		validationErrors = append(validationErrors, errors.New("password must be at least 8 characters long"))
	}

	return validationErrors
}
