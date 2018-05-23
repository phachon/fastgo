package main

import (
	"github.com/phachon/fastgo"
	"fastgo/_example/mvc/app/controllers"
)

func main() {

	// /author/login
	fastgo.Route.Add("GET", "/author/login", controllers.NewAuthor(), "Login")
	// /author/save
	fastgo.Route.POST("/author/save", controllers.NewAuthor(), "Save")
	// /main/index
	fastgo.Route.Get("/main/index", controllers.NewMain(), "Index")

	fastgo.SetStaticPath("static")
	fastgo.SetViewsPath("views")

	fastgo.Run()
}