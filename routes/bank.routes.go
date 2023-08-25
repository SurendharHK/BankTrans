package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

func BankingRoute(route *gin.Engine, controller controllers.BankingController) {
	route.POST("/banking/create", controller.CreateCustomer)
	route.GET("/customers", controller.GetCustomers)
	route.POST("/banking/updatecustomer", controller.UpdateCustomer)
	route.POST("/banking/deletecustomer", controller.DeleteCustomer)
	
}
