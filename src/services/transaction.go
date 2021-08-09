package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/middlewares"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/utils"
	"errors"
)

type TransactionService struct {}

func (TransactionService) GetUserTransactionList(userId, page, size int) (interface{}, error) {
	var count int64
	db := database.GetDB().
		Model(&models.Transaction{}).
		Where("user_id = ?", userId).
		Count(&count)

	transactionList := make([]models.Transaction, 0)
	pagination := utils.Pagination{}

	err := db.
		Scopes(pagination.Paginate(page, size, "transaction_dtm desc")).
		Find(&transactionList).Error
	if err != nil {
		return nil, err
	}

	paginationResponse := utils.SetResponse(&pagination, "transactions", transactionList, count)
	return paginationResponse, err
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

func (TransactionService) GetPopularSnackList(start string, end string, transactionTypeId int, limit int) (interface{}, error) {
	if !isValidQueryParams(start, end, transactionTypeId, limit) {
		return nil, errors.New(middlewares.BadRequest)
	}

	cond := "transaction_type_id = ? AND payment_id <> 0 AND transaction_dtm BETWEEN ? AND ?"
	db := database.GetDB().
		Table("transactions").
		Select("snack_name, sum(quantity) as total_quantity").
		Order("total_quantity desc").
		Where(cond, transactionTypeId, start, end).
		Group("snack_name")

	popularSnacks := make([]models.PopularSnack, 0)
	err := db.Limit(limit).Find(&popularSnacks).Error
	if err != nil {
		return nil, err
	}

	for i := range popularSnacks {
		snack := models.Snack{}
		db = database.GetDB().Model(&models.Snack{})
		err = db.Find(&snack, models.Snack{Name: popularSnacks[i].SnackName}).Error
		if err != nil {
			return nil, errors.New(middlewares.NotFound)
		}
		popularSnacks[i].SnackTypeID = snack.SnackTypeID
		popularSnacks[i].ImageURI = snack.ImageURI
	}

	return popularSnacks, nil
}

func isValidQueryParams(start string, end string, transactionTypeId int, limit int) bool {
	return !(start == "" || end == "" || transactionTypeId == 0 || limit == 0)
}
