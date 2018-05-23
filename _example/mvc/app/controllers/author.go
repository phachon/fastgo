package controllers

import (
	"strings"
	"github.com/phachon/fastgo"
	"github.com/phachon/fastgo/_example/mvc/app/utils"
)

func NewAuthor() *AuthorController {
	return &AuthorController{}
}

type AuthorController struct {
	BaseController
}

// login index
func (this *AuthorController) Index() {

	this.Data["title"] = "login"
	this.LayoutRender("layouts/author", "author/login")
}

// login
func (this *AuthorController) Login() {

	username := strings.TrimSpace(this.GetString("username", ""))
	password := strings.TrimSpace(this.GetString("password", ""))

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

	if username != realUser || password != realPass {
		this.jsonError("username or password error!")
		return
	}

	user := map[string]string{
		"username": username,
		"password": password,
	}

	// save session
	this.Session.Set("author", user)
	// save cookie
	identify := utils.Encrypt.Md5Encode(string(this.UserAgent()) + this.Ctx.RemoteIP().String() + password)
	passportValue := utils.Encrypt.Base64Encode(username + "@" + identify)
	passport := fastgo.Conf.GetString("author.passport")
	this.SetCookie(passport, passportValue, 3600)

	this.jsonSuccess("login success", nil, "/main/index")
}

// logout
func (this *AuthorController) Logout() {

	this.Session.Delete("author")
	this.Ctx.Redirect("/", 302)
}