package handlers

import (
	"Snack-Golang-Server/src/services"
	"Snack-Golang-Server/src/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SnackList(c *gin.Context) {
	snackService := services.SnackService{}
	snackList, err := snackService.GetSnackList()
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"snacks": snackList,
	})
}

func SnackCreate(c *gin.Context) {
	snackService := services.SnackService{}

	snackRegisterRequest := validators.SnackRegisterRequest{}
	if err := c.ShouldBindJSON(&snackRegisterRequest); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	res, err := snackService.AddSnack(snackRegisterRequest)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func SnackUpdate(c *gin.Context) {

}

func SnackDelete(c *gin.Context) {

}
