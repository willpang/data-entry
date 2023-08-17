package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/willpang/data-entry/models"
)

func LoadIndexTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"Title":   "Home Page",
		"Message": "Hello, World!",
	})
}

func LoadUserTemplate(c *gin.Context) {

	c.HTML(http.StatusOK, "users/users.tmpl", gin.H{
		"Title": "Users Page",
		"UserData": []models.User{
			{
				Name:        "Test 1",
				Address:     "Street 1",
				Email:       "Test1@gmail.com",
				PhoneNumber: "0811234567890",
			},
			{
				Name:        "Test 2",
				Address:     "Street 2",
				Email:       "Test2@gmail.com",
				PhoneNumber: "0821234567890",
			},
		},
	})
}
