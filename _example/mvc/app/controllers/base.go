package controllers

import (
	"github.com/phachon/fastgo"
	"strings"
	"github.com/phachon/fastgo/_example/mvc/app/utils"
)

func NewBaseController() *BaseController {
	return &BaseController{}
}

type BaseController struct {
	fastgo.Controller
	LoginUser map[string]string
	controller string
}

type JsonResult struct {
	Code int `json:"code"`
	Message interface{} `json:"message"`
	Data interface{} `json:"data"`
	Redirect map[string]interface{} `json:"redirect"`
}

func (this *BaseController) Before() {
	this.controller = strings.ToLower(this.ControllerName[0 : len(this.ControllerName)-10])
	if this.controller != "author" && !this.isLogin() {
		this.Ctx.Redirect("/author/index", 302)
		return
	}
}

func (this *BaseController) isLogin() bool {
	passport := fastgo.Conf.GetString("author.passport")

	cookie := this.GetCookie(passport)
	// cookie is empty
	if len(cookie) == 0 {
		return false
	}
	user := this.Session.Get("author")
	// session is empty
	if user == nil {
		return false
	}
	cookieValue, _ := utils.Encrypt.Base64Decode(string(cookie))
	identifyList := strings.Split(cookieValue, "@")
	if cookieValue == "" || len(identifyList) != 2 {
		return false
	}
	username := identifyList[0]
	identify := identifyList[1]
	userValue := user.(map[string]string)

	// cookie session name
	if username != userValue["username"] {
		return false
	}
	passStr := string(this.UserAgent()) + this.Ctx.RemoteIP().String() + userValue["password"]
	// UAG and IP
	if identify != utils.Encrypt.Md5Encode(passStr) {
		return false
	}
	this.LoginUser = userValue
	// success
	return true
}

// return json error
func (this *BaseController) jsonError(message interface{}, data ...interface{}) {
	this.jsonResult(0, message, data...)
}

// return json success
func (this *BaseController) jsonSuccess(message interface{}, data ...interface{}) {
	this.jsonResult(1, message, data...)
}

// return json result
func (this *BaseController) jsonResult(code int, message interface{}, data ...interface{}) {

	if message == nil {
		message = ""
	}
	url := ""
	sleep := 2000
	var _data interface{}
	if len(data) > 0 {
		_data = data[0]
	}
	if len(data) > 1 {
		url = data[1].(string)
	}
	if len(data) > 2 {
		sleep = data[2].(int)
	}
	res := JsonResult {
		Code:    code,
		Message: message,
		Data:    _data,
		Redirect: map[string]interface{}{
			"url":   url,
			"sleep": sleep,
		},
	}
	this.ReturnJson(res)
}

func (this *BaseController) viewError(message string) {
	this.Render("error/error")
}