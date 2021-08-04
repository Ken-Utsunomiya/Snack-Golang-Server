package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
)

type TransactionService struct {}

func (TransactionService) GetUserTransactionList(userId int) ([]models.Transaction, error) {
	db := database.GetDB()
	transactions := make([]models.Transaction, 0)
	err := db.Find(&transactions, models.Transaction{UserID: userId}).Error
	return transactions, err
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
