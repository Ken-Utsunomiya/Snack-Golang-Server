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

	userId, _ := strconv.Atoi(c.Param("user_id"))

	_, err := userService.GetUser(userId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	userPayment, err := paymentService.GetUserPaymentList(userId, page, size)

	c.JSON(http.StatusOK, userPayment)
}

func PaymentCreate(c *gin.Context) {

}

func PaymentAllCreate(c *gin.Context) {

}
