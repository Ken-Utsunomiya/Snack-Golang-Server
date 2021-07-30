package routers

import (
	"github.com/gin-gonic/gin"

	"Snack-Golang-Server/src/handlers"
)

func TransactionRoutes(rg *gin.RouterGroup) {

	transaction := rg.Group("/transactions")
	{
		transaction.GET("/", handlers.PopularSnackList)
		transaction.POST("/", handlers.TransactionCreate)
		transaction.PUT("/:transaction_id", handlers.TransactionUpdate)
	}
}