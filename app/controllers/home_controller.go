package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
}

func (controller HomeController) Index(context *gin.Context) {
	context.Query("id")
	context.HTML(http.StatusOK, "index.go.html", gin.H{
		"name": "Dunia",
	})
}
