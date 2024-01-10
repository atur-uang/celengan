package controllers

import (
	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
}

func (controller HomeController) Index(context *gin.Context) {
	context.Query("id")
	context.HTML(http.StatusOK, "home/index.html", pongo2.Context{
		"name": "Dunia",
	})
}

func (controller HomeController) Hello(context *gin.Context) {
	context.JSON(http.StatusOK, map[string]string{
		"hello": "world",
	})
}
