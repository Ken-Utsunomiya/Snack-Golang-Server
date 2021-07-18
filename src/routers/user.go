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
		user.GET("/:user_id/transactions")
		user.GET("/:user_id/transactions/:transaction_id")
		user.GET("/:user_id/payments")
		user.GET("/:user_id/pendingOrders")
		user.POST("/", handlers.AddUser)
		user.PUT("/:user_id", handlers.UpdateUser)
		user.DELETE("/:user_id", handlers.DeleteUser)
	}
}
