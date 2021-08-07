package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/utils"
	"math"
)

type PaymentService struct {}

func (PaymentService) GetUserPaymentList(userId, page, size int) (utils.PaginationResponse, error)  {
	var count int64
	db := database.GetDB().Model(&models.Payment{}).Where("user_id = ?", userId).Count(&count)
	paymentList := make([]models.Payment, 0)
	pagination := utils.Pagination{}
	paginationResponse := utils.PaginationResponse{}
	err := db.Scopes(pagination.Paginate(db, page, size, "payment_dtm desc")).Find(&paymentList).Error
	if err != nil {
		return paginationResponse, err
	}
	paginationResponse.TotalRows = int(count)
	paginationResponse.TotalPages = int(math.Ceil(float64(paginationResponse.TotalRows / pagination.Limit)))
	paginationResponse.Payments = paymentList
	paginationResponse.CurrentPage = page
	return paginationResponse, err
}

func (PaymentService) AddPayment(payment *models.Payment) error {
	return nil
}

func (PaymentService) AddPaymentAll(payments []models.Payment) error {
	return nil
}
