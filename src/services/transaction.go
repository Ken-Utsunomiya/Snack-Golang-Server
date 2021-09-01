package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/middlewares"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/utils"
	"Snack-Golang-Server/src/validators"
	"errors"
	"gorm.io/gorm"
)
const (
	PURCHASE = 1
	CANCEL = 2
	PENDING = 3
	PENDING_CANCEL = 4
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

func (TransactionService) AddTransaction(request validators.TransactionRegisterRequest) (models.Transaction, error) {
	db := database.GetDB()

	var transaction models.Transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		user := models.User{}
		userId := request.UserID
		if err := tx.First(&user, userId).Error; err != nil {
			return err
		}

		snack := models.Snack{}
		snackId := request.SnackID
		if err := tx.First(&snack, snackId).Error; err != nil {
			return err
		}

		// Update snack batch here

		if request.TransactionTypeID == PURCHASE {
			balance := user.Balance + request.TransactionAmount
			if err := tx.Model(&user).Update("balance", balance).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		transaction = validators.RegisterRequestToTransactionModel(request, snack.Name)

		if err := tx.Model(models.Transaction{}).Create(&transaction).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return transaction, err
}

func (TransactionService) UpdateTransaction(request validators.TransactionUpdateRequest, id int) (models.Transaction, error) {
	db := database.GetDB()

	var transaction models.Transaction
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&transaction, id).Error; err != nil {
			return err
		}

		userId := transaction.UserID
		//if !utils.IsValidUser(userId) {
		//	return errors.New(middlewares.NotAuthorized)
		//}
		user := models.User{}
		if err := tx.First(&user, userId).Error; err != nil {
			return err
		}

		snack := models.Snack{}
		snackName := transaction.SnackName
		if err := tx.First(&snack, models.Snack{Name: snackName}).Error; err != nil {
			return err
		}

		from := transaction.TransactionTypeID
		to := request.TransactionTypeID
		amount := transaction.TransactionAmount

		if from == to {
			return nil
		} else if from == PURCHASE && to == CANCEL {
			if transaction.PaymentID != nil {
				return errors.New(middlewares.BadRequest)
			}
			// increase snackbatch quantity
			// user balance decrement
			user.Balance -= amount
			if err := tx.Save(&user).Error; err != nil {
				return err
			}
		} else if from == PENDING && to == PURCHASE {
			// user balance increment
			user.Balance += amount
			if err := tx.Save(&user).Error; err != nil {
				return err
			}
		} else if from == PENDING && to == PENDING_CANCEL {
			if transaction.PaymentID != nil {
				return errors.New(middlewares.BadRequest)
			}
			// increase snackbatch quantity
		} else {
			return errors.New(middlewares.BadRequest)
		}

		transaction.TransactionTypeID = to
		if err := tx.Save(&transaction).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return transaction, err
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

func (TransactionService) GetPopularSnackList(start string, end string, transactionTypeId int, limit int) ([]models.PopularSnack, error) {
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
		err = db.First(&snack, models.Snack{Name: popularSnacks[i].SnackName}).Error
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
