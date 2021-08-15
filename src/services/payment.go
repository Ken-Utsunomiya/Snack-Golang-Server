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

type PaymentService struct {}

func (PaymentService) GetUserPaymentList(userId, page, size int) (interface{}, error)  {
	var count int64
	db := database.GetDB().
		Model(&models.Payment{}).
		Where("user_id = ?", userId).
		Count(&count)

	paymentList := make([]models.Payment, 0)
	pagination := utils.Pagination{}

	err := db.
		Scopes(pagination.Paginate(page, size, "payment_dtm desc")).
		Find(&paymentList).Error
	if err != nil {
		return nil, err
	}

	paginationResponse := utils.SetResponse(&pagination, "payments", paymentList, count)
	return paginationResponse, err
}

func (PaymentService) AddPayment(request validators.PaymentRegisterRequest) (models.Payment, error) {
	db := database.GetDB()

	var payment models.Payment
	err := db.Transaction(func(tx *gorm.DB) error {
		user := models.User{}
		userId := request.UserID
		if err := tx.First(&user, userId).Error; err != nil {
			return err
		}

		// decrease user balance
		updatedBalance := user.Balance - request.PaymentAmount
		if updatedBalance < 0 {
			return errors.New(middlewares.BadRequest)
		}

		if err := tx.Model(&user).Update("balance", updatedBalance).Error; err != nil {
			return err
		}

		// create a payment


		return nil
	})
	return payment, err
}

func (PaymentService) AddPaymentAll(payments []models.Payment) error {
	return nil
}
