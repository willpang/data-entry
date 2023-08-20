package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/datastore/sqlite"
	"github.com/willpang/data-entry/handler"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.tmpl")

	sqlite.ConnectDatabase()

	r.GET("/", handler.LoadIndexTemplate)
	r.GET("/users", handler.LoadUserTemplate)
	r.GET("/user", handler.FindUsers)
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.FindUser)
	r.PATCH("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	r.Run()
}
