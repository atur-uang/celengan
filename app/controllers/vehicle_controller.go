package controllers

import (
	"github.com/atur-uang/celengan/app/models"
	"github.com/atur-uang/celengan/framework"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type VehicleController struct {
}

func (controller VehicleController) Index(context *gin.Context) {
	var response []models.VehicleResponse
	db := framework.GetDB()
	db.Model(&models.Vehicle{}).Find(&response)

	context.JSON(http.StatusOK, gin.H{"data": response})

}

func (controller VehicleController) Store(context *gin.Context) {
	var vehicle models.Vehicle

	if JSONError := context.ShouldBindJSON(&vehicle); JSONError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": JSONError.Error()})
		return
	}
	db := framework.GetDB()

	result := db.Create(&vehicle)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create the vehicle"})
		return
	}

	response := models.VehicleResponse{
		ID:   vehicle.ID,
		Name: vehicle.Name,
	}

	context.JSON(http.StatusCreated, gin.H{"data": response})
}

func (controller VehicleController) Detail(context *gin.Context) {
	id, parseError := strconv.ParseUint(context.Param("id"), 10, 64)
	if parseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	var vehicle models.VehicleResponse
	db := framework.GetDB()
	db.Model(&models.Vehicle{}).First(&vehicle, uint(id))
	context.JSON(http.StatusOK, gin.H{"data": vehicle})
}

func (controller VehicleController) Update(context *gin.Context) {
	id, parseError := strconv.ParseUint(context.Param("id"), 10, 64)
	if parseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	var existingVehicle models.Vehicle
	db := framework.GetDB()
	result := db.First(&existingVehicle, uint(id))
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Failed to get the existing vehicle"})
		return
	}

	var updatedVehicle models.Vehicle
	if JSONError := context.ShouldBindJSON(&updatedVehicle); JSONError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": JSONError.Error()})
		return
	}

	result = db.Model(&existingVehicle).Updates(&updatedVehicle)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update the vehicle"})
		return
	}

	response := models.VehicleResponse{
		ID:   existingVehicle.ID,
		Name: updatedVehicle.Name,
	}
	context.JSON(http.StatusOK, gin.H{"data": response})

}

func (controller VehicleController) Delete(context *gin.Context) {
	// Get the vehicle ID from the URL parameter
	id, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	db := framework.GetDB()

	// Check if the vehicle with the given ID exists
	var vehicle models.Vehicle
	result := db.First(&vehicle, uint(id))
	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	// Delete the vehicle from the database
	result = db.Delete(&vehicle)

	// Check for errors
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the vehicle"})
		return
	}

	// Return success message
	context.String(http.StatusNoContent, "")
}
