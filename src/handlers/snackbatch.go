package handlers

import (
	"Snack-Golang-Server/src/services"
	"Snack-Golang-Server/src/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SnackBatchList(c *gin.Context) {
	snackId, _ := strconv.Atoi(c.Query("snack_id"))

	var order string
	if snackId == 0 {
		order = "snack_batch_id"
	} else {
		order = "expiration_dtm, snack_batch_id"
	}

	snackBatchService := services.SnackBatchService{}
	snackBatchList, err := snackBatchService.GetSnackBatchList(snackId, order)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{ "snack_batches": snackBatchList })
}

func SnackBatchCreate(c *gin.Context) {
	snackBatchService := services.SnackBatchService{}

	snackBatchRegisterRequest := validators.SnackBatchRegisterRequest{}
	if err := c.ShouldBindJSON(&snackBatchRegisterRequest); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	res, err := snackBatchService.AddSnackBatch(snackBatchRegisterRequest)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func SnackBatchUpdate(c *gin.Context) {

}

func SnackBatchDelete(c *gin.Context) {
	snackBatchService := services.SnackBatchService{}

	snackBatchId, _ := strconv.Atoi(c.Param("snack_batch_id"))
	if err := snackBatchService.DeleteSnackBatch(snackBatchId); err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusNoContent)
}
