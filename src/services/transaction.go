package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/utils"
)

type TransactionService struct {}

func (TransactionService) GetUserTransactionList(userId int) ([]models.Transaction, error) {
	db := database.GetDB()
	transactions := make([]models.Transaction, 0)
	err := db.Find(&transactions, models.Transaction{UserID: userId}).Error
	return transactions, err
}

func (TransactionService) GetUserTransaction(userId int, transactionId int) (models.Transaction, error) {
	db := database.GetDB()
	transaction := models.Transaction{}
	err := db.First(
		&transaction,
		models.Transaction{
			ID: transactionId,
			UserID: userId,
		}).Error
	return transaction, err
}

func (TransactionService) AddTransaction(transaction *models.Transaction) error {
	return nil
}

func (TransactionService) UpdateTransaction(transaction *models.Transaction) error {
	return nil
}

func (TransactionService) GetPendingOrderList(userId, page, size int) (interface{}, error) {
	var count int64
	db := database.GetDB().
		Model(&models.Transaction{}).
		Where("user_id = ? AND transaction_type_id = ?", userId, 3).
		Count(&count)

	pendingOrderList := make([]models.Transaction, 0)
	pagination := utils.Pagination{}

	err := db.
		Scopes(pagination.Paginate(page, size, "transaction_dtm desc")).
		Find(&pendingOrderList).Error
	if err != nil {
		return nil, err
	}

	paginationResponse := utils.SetResponse(&pagination, "transactions", pendingOrderList, count)
	return paginationResponse, err
}

func (TransactionService) GetPopularSnackList() []models.Snack {
	return nil
}
