package handlers

import (
	"Snack-Golang-Server/src/services"
	"Snack-Golang-Server/src/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	snackService := services.SnackService{}

	snackUpdateRequest := validators.SnackUpdateRequest{}
	if err := c.ShouldBindJSON(&snackUpdateRequest); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	snackId, _ := strconv.Atoi(c.Param("snack_id"))

	res, err := snackService.UpdateSnack(snackUpdateRequest, snackId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func SnackDelete(c *gin.Context) {
	snackService := services.SnackService{}

	snackId, _ := strconv.Atoi(c.Param("snack_id"))
	if err := snackService.DeleteSnack(snackId); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusNoContent)
}
