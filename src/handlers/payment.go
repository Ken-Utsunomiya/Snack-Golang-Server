package handlers

import (
	"Snack-Golang-Server/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserPaymentList(c *gin.Context) {
	userService := services.UserService{}
	paymentService := services.PaymentService{}

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

	userPaymentList, err := paymentService.GetUserPaymentList(userId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, userPaymentList)
}

func PaymentCreate(c *gin.Context) {

}

func PaymentAllCreate(c *gin.Context) {

}
