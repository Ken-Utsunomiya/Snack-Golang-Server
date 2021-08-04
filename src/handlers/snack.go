package handlers

import (
	"Snack-Golang-Server/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SnackList(c *gin.Context) {
	snackService := services.SnackService{}
	snackList, err := snackService.GetSnackList()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"snacks": snackList,
	})
}

func SnackCreate(c *gin.Context) {

}

func SnackUpdate(c *gin.Context) {

}

func SnackDelete(c *gin.Context) {

}
