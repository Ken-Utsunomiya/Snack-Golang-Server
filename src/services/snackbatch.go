package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
)

type SnackBatchService struct {}

func (SnackBatchService) GetSnackBatchList(id int, order string) ([]models.SnackBatch, error) {
	db := database.GetDB()
	snackBatches := make([]models.SnackBatch, 0)

	var err error
	if id == 0 {
		err = db.
			Order(order).
			Find(&snackBatches).Error
	} else {
		err = db.
			Order(order).
			Find(&snackBatches, models.SnackBatch{SnackID: id}).Error
	}

	return snackBatches, err
}

func (SnackBatchService) AddSnackBatch(snackBatch *models.SnackBatch) error {
	return nil
}

func (SnackBatchService) UpdateSnackBatch(snackBatch *models.SnackBatch) error {
	return nil
}

func (SnackBatchService) DeleteSnackBatch(id int) error {
	db := database.GetDB()
	err := db.Delete(&models.SnackBatch{}, id).Error
	return err
}
