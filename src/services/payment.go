package services

import "Snack-Golang-Server/src/models"

type PaymentService struct {}

func (PaymentService) GetUserPaymentList() []models.Payment {
	return nil
}

func (PaymentService) AddPayment(payment *models.Payment) error {
	return nil
}

func (PaymentService) AddPaymentAll(payments []models.Payment) error {
	return nil
}
