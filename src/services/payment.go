package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/utils"
)

type PaymentService struct {}

func (PaymentService) GetUserPaymentList(userId, page, size int) (utils.PaginationResponse, error)  {
	var count int64
	db := database.GetDB().Model(&models.Payment{}).Where("user_id = ?", userId).Count(&count)

	paymentList := make([]models.Payment, 0)
	pagination := utils.Pagination{}
	paginationResponse := utils.PaginationResponse{}

	err := db.Scopes(pagination.Paginate(page, size, "payment_dtm desc")).Find(&paymentList).Error
	if err != nil {
		return paginationResponse, err
	}

	utils.SetResponse(&pagination, &paginationResponse, paymentList, count)
	return paginationResponse, err
}

func (PaymentService) AddPayment(payment *models.Payment) error {
	return nil
}

func (PaymentService) AddPaymentAll(payments []models.Payment) error {
	return nil
}
