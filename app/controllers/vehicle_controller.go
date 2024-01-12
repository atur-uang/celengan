package controllers

import (
	"github.com/atur-uang/celengan/app/models"
	"github.com/atur-uang/celengan/framework"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VehicleController struct {
}

func (controller VehicleController) Index(context *gin.Context) {
	var vehicles models.VehicleAPI
	db := framework.GetDB()
	db.Model(&models.Vehicle{}).Find(vehicles)

	context.JSON(http.StatusOK, gin.H{"data": vehicles})

}
