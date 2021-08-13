package validators

import "Snack-Golang-Server/src/models"

type TransactionRegisterRequest struct {
	UserID						int `json:"user_id" binding:"required"`
	TransactionTypeID	int `json:"transaction_type_id" binding:"required"`
	SnackID						int `json:"snack_id" binding:"required"`
	TransactionAmount	int `json:"transaction_amount" binding:"required,min=0"`
	Quantity 					int `json:"quantity" binding:"required,min=1"`
}

func RegisterRequestToTransactionModel(request TransactionRegisterRequest, snackName string) models.Transaction {
	transaction := models.Transaction{}
	transaction.UserID = request.UserID
	transaction.TransactionTypeID = request.TransactionTypeID
	transaction.SnackName = snackName
	transaction.TransactionAmount = request.TransactionAmount
	transaction.Quantity = request.Quantity
	return transaction
}
