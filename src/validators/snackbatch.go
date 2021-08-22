package validators

import (
	"Snack-Golang-Server/src/models"
	"time"
)

type SnackBatchRegisterRequest struct {
	SnackID	int `json:"snack_id" binding:"required"`
	Quantity	int `json:"quantity" binding:"required,min=1"`
	ExpirationDTM	*time.Time `json:"expiration_dtm" binding:"required"`
}

type SnackBatchUpdateRequest struct {
	Quantity	int	`json:"quantity" binding:"min=0"`
	ExpirationDTM	*time.Time	`json:"expiration_dtm"`
}

func RegisterRequestToSnackBatchModel(request SnackBatchRegisterRequest) models.SnackBatch {
	snackbatch := models.SnackBatch{}
	snackbatch.SnackID = request.SnackID
	snackbatch.Quantity = request.Quantity
	snackbatch.ExpirationDTM = request.ExpirationDTM
	return snackbatch
}

func UpdateRequestToSnackBatchModel(request SnackBatchUpdateRequest, snackbatch *models.SnackBatch) {
	snackbatch.Quantity = request.Quantity
	if request.ExpirationDTM != nil {
		snackbatch.ExpirationDTM = request.ExpirationDTM
	}
}
