package app

import (
	"github.com/atur-uang/celengan/app/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) *gin.Engine {

	//route.GET("/", controllers.HomeController{}.Index)
	route.GET("hello", controllers.HomeController{}.Hello)
	route.GET("hello/world", controllers.HomeController{}.Hello)
	route.POST("transaction", controllers.TransactionController{}.Deposit)
	route.DELETE("transaction", controllers.TransactionController{}.Withdrawal)

	//vehicles := route.Group("vehicles")
	//vehicles.GET("/", (&controllers.VehicleController{}).Index)
	//vehicles.POST("/", (&controllers.VehicleController{}).Store)

	route.GET("vehicles", controllers.VehicleController{}.Index)
	route.POST("vehicles", controllers.VehicleController{}.Store)
	route.GET("vehicles/:id", controllers.VehicleController{}.Detail)
	route.PUT("vehicles/:id", controllers.VehicleController{}.Update)
	route.DELETE("vehicles/:id", controllers.VehicleController{}.Delete)

	return route
}
