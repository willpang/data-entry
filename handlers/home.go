package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) LoadIndexTemplate(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home/index.tmpl", gin.H{
		"Page":    "Home",
		"Title":   "Home",
		"Message": "Hello, World!",
	})
}
