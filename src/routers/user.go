package routers

import (
	"github.com/gin-gonic/gin"

	"Snack-Golang-Server/src/handlers"
)

func UserRoutes(rg *gin.RouterGroup) {

	user := rg.Group("/users")
	{
		user.GET("/", handlers.GetUsers)
		user.GET("/common", handlers.GetUsersCommon)
		user.GET("/:user_id", handlers.GetUser)
		user.GET("/:user_id/transactions", handlers.GetUserTransactions)
		user.GET("/:user_id/transactions/:transaction_id", handlers.GetUserTransaction)
		user.GET("/:user_id/payments", handlers.UserPaymentList)
		user.GET("/:user_id/pendingOrders", handlers.GetPendingOrders)
		user.POST("/", handlers.AddUser)
		user.PUT("/:user_id", handlers.UpdateUser)
		user.DELETE("/:user_id", handlers.DeleteUser)
	}
}
