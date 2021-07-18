package routers

import (
	"Snack-Golang-Server/src/handlers"
	"github.com/gin-gonic/gin"
)

func SnackBatchRoutes(rg *gin.RouterGroup) {

	snackBatch := rg.Group("/snackBatches")
	{
		snackBatch.GET("/", handlers.GetSnackBatches)
		snackBatch.POST("/", handlers.AddSnackBatch)
		snackBatch.PUT("/:snackBatchId", handlers.UpdateSnackBatch)
		snackBatch.DELETE("/:snackBatchId", handlers.DeleteSnackBatch)
	}
}
