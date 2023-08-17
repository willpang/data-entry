package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadIndexTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":   "Main website",
		"message": "Hello, World!",
	})
}
