package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/utils"
	"Snack-Golang-Server/src/validators"
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
	return models.Payment{}, nil
}

func (PaymentService) AddPaymentAll(payments []models.Payment) error {
	return nil
}
