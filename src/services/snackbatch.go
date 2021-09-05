package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/middlewares"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/validators"
	"gorm.io/gorm"
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

	// create a new snackbatch
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

	// convert update request to snackbatch model
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

func (SnackBatchService) IncreaseQuantityInSnackBatch(snackId int, quantity int, tx *gorm.DB) error {
	snackbatch := models.SnackBatch{}

	err := tx.Order("expiration_dtm, snack_batch_id").First(&snackbatch, models.SnackBatch{SnackID: snackId}).Error
	if err.Error() == middlewares.RecordNotFound {
		request := validators.SnackBatchRegisterRequest{SnackID: snackId, Quantity: quantity, ExpirationDTM: nil}
		snackbatch = validators.RegisterRequestToSnackBatchModel(request)
		if err := tx.Create(&snackbatch).Error; err != nil {
			return err
		}
		return nil
	} else if snackbatch != (models.SnackBatch{}) {
		snackbatch.Quantity += quantity
		if err := tx.Save(&snackbatch).Error; err != nil {
			return err
		}
		return nil
	}

	return err
}
