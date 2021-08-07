package routers

import (
	"Snack-Golang-Server/src/handlers"
	"github.com/gin-gonic/gin"
)

func SnackBatchRoutes(rg *gin.RouterGroup) {

	snackBatch := rg.Group("/snack_batches")
	{
		snackBatch.GET("/", handlers.SnackBatchList)
		snackBatch.POST("/", handlers.SnackBatchCreate)
		snackBatch.PUT("/:snackBatchId", handlers.SnackBatchUpdate)
		snackBatch.DELETE("/:snackBatchId", handlers.SnackBatchDelete)
	}
}
