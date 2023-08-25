package routes

import (
	"github.com/gin-gonic/gin"
	
	"main.go/controllers"
)

func TransactionRoutes(route *gin.Engine,controller controllers.TransactionController){
	route.POST("/banking/transfer",controller.Transfer)
	
}