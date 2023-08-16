package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/controllers"
	"github.com/willpang/data-entry/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)

	r.Run()
}
