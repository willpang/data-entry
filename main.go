package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/datastore/sqlite"
	"github.com/willpang/data-entry/handler"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.tmpl")

	db := sqlite.ConnectDatabase()

	h := &handler.Handler{
		DB: db,
	}

	r.GET("/", h.LoadIndexTemplate)
	r.GET("/users", h.LoadUserTemplate)
	r.GET("/user", h.FindUsers)
	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.FindUser)
	r.PATCH("/user/:id", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)

	r.Run()
}
