package main

import (
	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/datastores/services"
	"github.com/willpang/data-entry/datastores/sqlite"
	"github.com/willpang/data-entry/handlers"
)

var (
	_ services.UserDatastore = &sqlite.UserDatastore{}
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*.tmpl")

	db := sqlite.ConnectDatabase()

	h := &handlers.Handler{
		UserDatastore: sqlite.UserDatastore{DB: db},
	}

	r.GET("/", h.LoadIndexTemplate)
	r.GET("/users", h.LoadUserTemplate)
	r.GET("/user", h.FindUsers)
	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.FindUser)
	r.PATCH("/user/:id", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)

	// Handle CSS, JS, and Image
	r.Static("/assets", "./assets/")

	r.Run()
}
