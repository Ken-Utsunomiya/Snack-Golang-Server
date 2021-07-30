package services

import "Snack-Golang-Server/src/models"

type TransactionService struct {}

func (TransactionService) GetUserTransactionList() []models.Transaction {
	return nil
}

func (TransactionService) GetUserTransaction(userId uint) *models.Transaction {
	return nil
}

func (TransactionService) AddTransaction(transaction *models.Transaction) error {
	return nil
}

func (TransactionService) UpdateTransaction(transaction *models.Transaction) error {
	return nil
}

func (TransactionService) GetPendingOrderList(userId uint) []models.Transaction {
	return nil
}

func (TransactionService) GetPopularSnackList() []models.Snack {
	return nil
}
