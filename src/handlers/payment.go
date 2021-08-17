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

func UserPaymentList(c *gin.Context) {
	userService := services.UserService{}
	paymentService := services.PaymentService{}

	userId, _ := strconv.Atoi(c.Param("user_id"))

	_, err := userService.GetUser(userId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	userPayment, err := paymentService.GetUserPaymentList(userId, page, size)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, userPayment)
}

func PaymentCreate(c *gin.Context) {
	paymentService := services.PaymentService{}

	registerRequest := validators.PaymentRegisterRequest{}
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.Error(errors.New(middlewares.BadRequest)).SetType(gin.ErrorTypePublic)
		return
	}

	//paymentUserId := registerRequest.UserID
	//
	//if !utils.IsValidUser(paymentUserId) {
	//	c.Error(errors.New(middlewares.NotAuthorized)).SetType(gin.ErrorTypePublic)
	//	return
	//}

	payment, err := paymentService.AddPayment(registerRequest)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, payment)
}

func PaymentAllCreate(c *gin.Context) {
	paymentService := services.PaymentService{}

	registerRequest := validators.PaymentRegisterRequest{}
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.Error(errors.New(middlewares.BadRequest)).SetType(gin.ErrorTypePublic)
		return
	}

	//paymentUserId := registerRequest.UserID
	//
	//if !utils.IsValidUser(paymentUserId) {
	//	c.Error(errors.New(middlewares.NotAuthorized)).SetType(gin.ErrorTypePublic)
	//	return
	//}

	payment, err := paymentService.AddPaymentAll(registerRequest)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, payment)
}
