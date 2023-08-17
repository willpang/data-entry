package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/controllers"
	"github.com/willpang/data-entry/models"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	models.ConnectDatabase()

	r.GET("/", controllers.LoadIndexTemplate)
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()
}
