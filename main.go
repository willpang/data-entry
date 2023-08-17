package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/controllers"
	"github.com/willpang/data-entry/models"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.tmpl")

	models.ConnectDatabase()

	r.GET("/", controllers.LoadIndexTemplate)
	r.GET("/users", controllers.LoadUserTemplate)
	r.GET("/user", controllers.FindUsers)
	r.POST("/user", controllers.CreateUser)
	r.GET("/user/:id", controllers.FindUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)

	r.Run()
}
