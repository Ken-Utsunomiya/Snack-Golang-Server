package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/validators"
	"gorm.io/gorm"
	"time"
)

type SnackService struct {}

func (SnackService) GetSnackList() ([]models.Snack, error) {
	db := database.GetDB()
	snacks := make([]models.Snack, 0)
	err := db.Find(&snacks).Error
	return snacks, err
}

func (SnackService) AddSnack(request validators.SnackRegisterRequest) (models.Snack, error) {
	db := database.GetDB()

	var snack models.Snack
	err := db.Transaction(func(tx *gorm.DB) error {
		snack = validators.RegisterRequestToSnackModel(request)

		if err := tx.Model(models.Snack{}).Create(snack).Error; err != nil {
			tx.Rollback()
			return err
		}

		if request.Quantity > 0 {
			snackId := snack.ID
			quantity := request.Quantity
			//var expirationDTM time.Time
			//if request.ExpirationDTM != nil {
			//	expirationDTM = *request.ExpirationDTM
			//} else {
			//	expirationDTM = nil
			//}
			snackbatch := models.SnackBatch{
				SnackID: snackId,
				Quantity: quantity,
				ExpirationDTM: nil,
			}

			if err := tx.Model(models.SnackBatch{}).Create(snackbatch).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

	}

	return snack, err
}

func (SnackService) UpdateSnack(snack *models.Snack) error {
	return nil
}

func (SnackService) DeleteSnack(id uint) error {
	return nil
}
