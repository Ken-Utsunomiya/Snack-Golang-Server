package routers

import (
	"Snack-Golang-Server/src/handlers"
	"github.com/gin-gonic/gin"
)

func SnackRoutes(rg *gin.RouterGroup) {

	snack := rg.Group("/snacks")
	{
		snack.GET("/", handlers.SnackList)
		snack.POST("/", handlers.SnackCreate)
		snack.PUT("/:snack_id", handlers.SnackUpdate)
		snack.DELETE("/:snack_id", handlers.SnackDelete)
	}
}
