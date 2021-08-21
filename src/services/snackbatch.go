package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/validators"
)

type SnackBatchService struct {}

func (SnackBatchService) GetSnackBatchList(snackId int, order string) ([]models.SnackBatch, error) {
	db := database.GetDB()
	snackBatches := make([]models.SnackBatch, 0)

	var err error
	if snackId == 0 {
		err = db.
			Order(order).
			Find(&snackBatches).Error
	} else {
		err = db.
			Order(order).
			Find(&snackBatches, models.SnackBatch{SnackID: snackId}).Error
	}

	return snackBatches, err
}

func (SnackBatchService) AddSnackBatch(request validators.SnackBatchRegisterRequest) (models.SnackBatch, error) {
	db := database.GetDB()

	snackbatch := validators.RegisterRequestToSnackBatchModel(request)

	err := db.Model(models.SnackBatch{}).Create(&snackbatch).Error
	return snackbatch, err
}

func (SnackBatchService) UpdateSnackBatch(request validators.SnackBatchUpdateRequest, id int) (models.SnackBatch, error) {
	db := database.GetDB()
	snackbatch := models.SnackBatch{}

	if err := db.First(&snackbatch, id).Error; err != nil {
		return snackbatch, err
	}

	if request == (validators.SnackBatchUpdateRequest{}) {
		return snackbatch, nil
	}

	validators.UpdateRequestToSnackBatchModel(request, &snackbatch)

	if err := db.Save(&snackbatch).Error; err != nil {
		return snackbatch, err
	}

	return snackbatch, nil
}

func (SnackBatchService) DeleteSnackBatch(id int) error {
	db := database.GetDB()
	err := db.Delete(&models.SnackBatch{}, id).Error
	return err
}
