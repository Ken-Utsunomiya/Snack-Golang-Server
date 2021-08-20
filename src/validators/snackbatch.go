package validators

import (
	"Snack-Golang-Server/src/models"
	"time"
)

type SnackBatchRegisterRequest struct {
	SnackID	int `json:"snack_id" binding:"required"`
	Quantity	int `json:"quantity" binding:"required,min=1"`
	ExpirationDTM	string `json:"expiration_dtm" binding:"required"`
}

type SnackBatchUpdateRequest struct {
	Quantity	int	`json:"quantity" binding:"required"`
	ExpirationDTM	*time.Time	`json:"expiration_dtm"`
}

func RegisterRequestToSnackBatchModel(request SnackBatchRegisterRequest) models.SnackBatch {
	snackbatch := models.SnackBatch{}
	snackbatch.SnackID = request.SnackID
	snackbatch.Quantity = request.Quantity
	//snackbatch.ExpirationDTM = utils.TimeStamp(request.ExpirationDTM)
	return snackbatch
}
