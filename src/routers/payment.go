package routers

import (
	"Snack-Golang-Server/src/handlers"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(rg *gin.RouterGroup) {

	payment := rg.Group("/payments")
	{
		payment.POST("/", handlers.AddPayment)
		payment.POST("/all", handlers.AddPaymentAll)
	}
}
