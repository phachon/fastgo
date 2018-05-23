package controllers

import (
	"github.com/phachon/fastgo"
)

func NewBaseController() *BaseController {
	return &BaseController{}
}

type BaseController struct {
	fastgo.Controller
	UserConf map[string]string
}

type JsonResult struct {
	Code int `json:"code"`
	Message interface{} `json:"message"`
	Data interface{} `json:"data"`
	Redirect map[string]interface{} `json:"redirect"`
}

func (b *BaseController) Before() {

}

// return json error
func (b *BaseController) jsonError(message interface{}, data ...interface{}) {
	b.jsonResult(0, message, data...)
}

// return json success
func (b *BaseController) jsonSuccess(message interface{}, data ...interface{}) {
	b.jsonResult(1, message, data...)
}

// return json result
func (b *BaseController) jsonResult(code int, message interface{}, data ...interface{}) {

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
	b.ReturnJson(res)
}

func (b *BaseController) viewError(message string) {
	b.Render("error/error")
}

func (b *BaseController) After() {

}