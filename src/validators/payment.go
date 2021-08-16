package validators

import (
	"Snack-Golang-Server/src/models"
	"time"
)

type PaymentRegisterRequest struct {
	UserID						int `json:"user_id" binding:"required"`
	CreatedBy	string `json:"created_by" binding:"required"`
	PaymentAmount	*int `json:"payment_amount" binding:"required"`
	TransactionIDs	[]int `json:"transaction_ids" binding:"required"`
}

func RegisterRequestToPaymentModel(request PaymentRegisterRequest) models.Payment {
	payment := models.Payment{}
	payment.UserID = request.UserID
	payment.CreatedBy = request.CreatedBy
	payment.PaymentAmount = *request.PaymentAmount
	payment.PaymentDTM = time.Now()
	return payment
}
