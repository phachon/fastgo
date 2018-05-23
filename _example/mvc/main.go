package main

import (
	"github.com/phachon/fastgo"
	"github.com/phachon/fastgo/_example/mvc/app/controllers"
)

func main() {

	author := controllers.NewAuthor()
	main := controllers.NewMain()

	// login
	fastgo.Route.GET("/author/index", author, "Index")
	fastgo.Route.POST("/author/login", author, "Login")
	fastgo.Route.GET("/author/logout", author, "Logout")

	// main
	fastgo.Route.GET("/", main, "Index")
	fastgo.Route.GET("/main/index", main, "Index")

	fastgo.Run()
}