package controllers

func NewMain() *MainController {
	return &MainController{}
}

type MainController struct {
	BaseController
}

func (this *MainController) Index() {

	//this.Data["user"] =
	this.Render("main/index")
}
