package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
)

type PaymentService struct {}

func (PaymentService) GetUserPaymentList(userId int) ([]models.Payment, error) {
	db := database.GetDB()
	paymentList := make([]models.Payment, 0)
	err := db.Order("payment_dtm desc").Find(&paymentList, models.Payment{ UserID: userId }).Error
	return paymentList, err
}

func (PaymentService) AddPayment(payment *models.Payment) error {
	return nil
}

func (PaymentService) AddPaymentAll(payments []models.Payment) error {
	return nil
}
