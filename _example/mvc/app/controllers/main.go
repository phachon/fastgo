package controllers

func NewMain() *MainController {
	return &MainController{}
}

type MainController struct {
	BaseController
}

func (this *MainController) Index() {

	this.Data["loginUser"] = this.LoginUser

	this.LayoutRender("layouts/default", "main/index")
}