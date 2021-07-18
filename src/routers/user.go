package routers

import (
	"github.com/gin-gonic/gin"

	"Snack-Golang-Server/src/handlers"
)

func UserRoutes(rg *gin.RouterGroup) {

	user := rg.Group("/users")
	{
		user.GET("/", handlers.GetUsers)
		//user.GET("/common")
		//user.GET("/:user_id")
		//user.GET("/:user_id/transactions")
		//user.GET("/:user_id/transactions/:transaction_id")
		//user.GET("/:user_id/payments")
		//user.GET("/:user_id/pendingOrders")
		//user.POST("/")
		//user.PUT("/:user_id")
		//user.DELETE("/:user_id")
	}
}
