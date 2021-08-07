package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
)

type SnackBatchService struct {}

func (SnackBatchService) GetSnackBatchList(Id int, order string) ([]models.SnackBatch, error) {
	db := database.GetDB()
	snackBatches := make([]models.SnackBatch, 0)

	var err error
	if Id == 0 {
		err = db.
			Order(order).
			Find(&snackBatches).Error
	} else {
		err = db.
			Order(order).
			Find(&snackBatches, models.SnackBatch{SnackID: Id}).Error
	}

	return snackBatches, err
}

func (SnackBatchService) AddSnackBatch(snackBatch *models.SnackBatch) error {
	return nil
}

func (SnackBatchService) UpdateSnackBatch(snackBatch *models.SnackBatch) error {
	return nil
}

func (SnackBatchService) DeleteSnackBatch(id uint) error {
	return nil
}
