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
		snackBatch.PUT("/:snack_batch_id", handlers.SnackBatchUpdate)
		snackBatch.DELETE("/:snack_batch_id", handlers.SnackBatchDelete)
	}
}
