package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"main.go/interfaces"
	"main.go/models"
)

type TransactionController struct {
	TransactionService interfaces.ITransaction
}

func InitTransactionController(profileService interfaces.ITransaction) TransactionController {
	return TransactionController{profileService}
}

func (t *TransactionController) Transfer(ctx *gin.Context) {
	var transactions models.Paying
	if err := ctx.ShouldBindJSON(&transactions); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := t.TransactionService.Transfer(transactions.From,transactions.To,transactions.Amount)

	if err != nil {
		if strings.Contains(err.Error(), "title already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newProfile})
}
