package handlers

import (
	"Snack-Golang-Server/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserTransactionList(c *gin.Context) {
	userService := services.UserService{}
	transactionService := services.TransactionService{}

	userId, _ := strconv.Atoi(c.Param("user_id"))

	_, err := userService.GetUser(userId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	userTransaction, err := transactionService.GetUserTransactionList(userId, page, size)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, userTransaction)
}

func UserTransactionRetrieve(c *gin.Context) {
	transactionService := services.TransactionService{}

	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	transactionId, err := strconv.Atoi(c.Param("transaction_id"))
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	userTransaction, err := transactionService.GetUserTransaction(userId, transactionId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, userTransaction)
}

func TransactionCreate(c *gin.Context) {

}

func TransactionUpdate(c *gin.Context) {

}

func PendingOrderList(c *gin.Context) {
	userService := services.UserService{}
	transactionService := services.TransactionService{}

	userId, _ := strconv.Atoi(c.Param("user_id"))
	_, err := userService.GetUser(userId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	pendingOrders, err := transactionService.GetPendingOrderList(userId, page, size)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, pendingOrders)
}

func PopularSnackList(c *gin.Context) {

}
