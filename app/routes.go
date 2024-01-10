package app

import (
	"github.com/atur-uang/celengan/app/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) *gin.Engine {

	//route.GET("/", controllers.HomeController{}.Index)
	route.GET("hello", controllers.HomeController{}.Hello)
	route.POST("transaction", controllers.TransactionController{}.Deposit)
	route.DELETE("transaction", controllers.TransactionController{}.Withdrawal)

	return route
}
