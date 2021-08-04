package handlers

import (
	"Snack-Golang-Server/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	userService = services.UserService{}
	transactionService = services.TransactionService{}
)

func UserTransactionList(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	_, err = userService.GetUser(userId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	userTransactionList, err := transactionService.GetUserTransactionList(userId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, userTransactionList)
}

func UserTransactionRetrieve(c *gin.Context) {

}

func TransactionCreate(c *gin.Context) {

}

func TransactionUpdate(c *gin.Context) {

}

func PendingOrderList(c *gin.Context) {

}

func PopularSnackList(c *gin.Context) {

}
