package routers

import (
	"Snack-Golang-Server/src/handlers"
	"github.com/gin-gonic/gin"
)

func SnackRoutes(rg *gin.RouterGroup) {

	snack := rg.Group("/snacks")
	{
		snack.GET("/", handlers.GetSnacks)
		snack.POST("/", handlers.AddSnack)
		snack.PUT("/:snack_id", handlers.UpdateSnack)
		snack.DELETE("/:snack_id", handlers.DeleteSnack)
	}
}
