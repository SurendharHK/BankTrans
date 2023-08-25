package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"main.go/interfaces"
	"main.go/models"
)

type BankingController struct {
	BankingService interfaces.IBank
}

func InitBankingController(profileService interfaces.IBank) BankingController {
	return BankingController{profileService}
}

func (pc *BankingController) CreateCustomer(ctx *gin.Context) {
	var profile *models.Account
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := pc.BankingService.CreateCustomer(profile)

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

func (pc *BankingController) GetCustomers(ctx *gin.Context) {

	customers, err := pc.BankingService.GetCustomers()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusFound, gin.H{"status": "success", "message": customers})

}

func (pc *BankingController) UpdateCustomer(ctx *gin.Context) {
	var profile *models.UpdateName
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	intialName := profile.IntialName
	updateName := profile.UpdateName

	customer, err := pc.BankingService.UpdateCustomer(intialName, updateName)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusFound, gin.H{"status": "success", "message": customer})

}

func (pc *BankingController) DeleteCustomer(ctx *gin.Context) {
	var name string
	if err := ctx.ShouldBindJSON(&name); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	customer, err := pc.BankingService.DeleteCustomer(name)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
	}
	ctx.JSON(http.StatusFound, gin.H{"status": "success", "message": customer})

}

func (pc *BankingController) Transfer(ctx *gin.Context) {

}

func (pc *BankingController) InsertManyCustomer(ctx *gin.Context) {
	var customers []interface{}
	if err := ctx.ShouldBindJSON(&customers); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newProfile, err := pc.BankingService.InsertManyCustomer(customers)

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
