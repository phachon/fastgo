package controllers

import (
	"strings"
	"github.com/phachon/fastgo"
)

func NewAuthor() *AuthorController {
	return &AuthorController{}
}

type AuthorController struct {
	BaseController
}

// login
func (this *AuthorController) Login() {

	this.Data["title"] = "login"
	this.LayoutRender("layout/login", "author/login")
}

// login
func (this *AuthorController) Save() {

	username := strings.Trim(this.GetString("username", ""), "")
	password := strings.Trim(this.GetString("password", ""), "")

	if username == "" {
		this.jsonError("username not empty!", nil)
		return
	}
	if password == "" {
		this.jsonError("password not empty!", nil)
		return
	}

	realUser := fastgo.Conf.GetString("admin.username")
	realPass := fastgo.Conf.GetString("admin.password")

	if username == realUser && password == realPass {
		this.jsonSuccess("login success", nil, "/main/index")
		return
	}

	this.jsonError("username or password error!")
}
