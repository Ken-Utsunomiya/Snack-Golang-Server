package services

import "Snack-Golang-Server/src/models"

type SnackBatchService struct {}

func (SnackBatchService) GetSnackBatchList() []models.SnackBatch {
	return nil
}

func (SnackBatchService) AddSnackBatch(snackBatch *models.SnackBatch) error {
	return nil
}

func (SnackBatchService) UpdateSnackBatch(snack *models.SnackBatch) error {
	return nil
}

func (SnackBatchService) DeleteSnackBatch(id uint) error {
	return nil
}
