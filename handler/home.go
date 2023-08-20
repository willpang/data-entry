package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadIndexTemplate(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"Title":   "Home Page",
		"Message": "Hello, World!",
	})
}
