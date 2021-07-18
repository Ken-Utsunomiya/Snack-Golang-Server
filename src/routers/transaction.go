package routers

import (
	"Snack-Golang-Server/src/handlers"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(rg *gin.RouterGroup) {

	transaction := rg.Group("/transactions")
	{
		transaction.GET("/", handlers.GetPopularSnacks)
		transaction.POST("/", handlers.AddTransaction)
		transaction.PUT("/:transaction_id", handlers.UpdateTransaction)
	}
}