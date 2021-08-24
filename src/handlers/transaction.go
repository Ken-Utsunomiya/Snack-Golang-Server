package handlers

import (
	"Snack-Golang-Server/src/middlewares"
	"Snack-Golang-Server/src/services"
	"Snack-Golang-Server/src/validators"
	"errors"
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
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	userTransaction, err := transactionService.GetUserTransactionList(userId, page, size)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, userTransaction)
}

func UserTransactionRetrieve(c *gin.Context) {
	transactionService := services.TransactionService{}

	userId, _ := strconv.Atoi(c.Param("user_id"))
	transactionId, _ := strconv.Atoi(c.Param("transaction_id"))

	userTransaction, err := transactionService.GetUserTransaction(userId, transactionId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, userTransaction)
}

func TransactionCreate(c *gin.Context) {
	transactionService := services.TransactionService{}

	registerRequest := validators.TransactionRegisterRequest{}
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		_ = c.Error(errors.New(middlewares.BadRequest)).SetType(gin.ErrorTypePublic)
		return
	}

	res, err := transactionService.AddTransaction(registerRequest)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func TransactionUpdate(c *gin.Context) {

}

func PendingOrderList(c *gin.Context) {
	userService := services.UserService{}
	transactionService := services.TransactionService{}

	userId, _ := strconv.Atoi(c.Param("user_id"))
	_, err := userService.GetUser(userId)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	pendingOrders, err := transactionService.GetPendingOrderList(userId, page, size)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, pendingOrders)
}

func PopularSnackList(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")
	transactionTypeId, _ := strconv.Atoi(c.Query("transaction_type_id"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	transactionService := services.TransactionService{}
	popularSnacks, err := transactionService.GetPopularSnackList(startDate, endDate, transactionTypeId, limit)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, popularSnacks)
}
