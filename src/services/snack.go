package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/validators"
	"gorm.io/gorm"
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
		if err := tx.Model(models.Snack{}).Create(&snack).Error; err != nil {
			tx.Rollback()
			return err
		}

		if request.Quantity > 0 {
			snackId := snack.ID
			quantity := request.Quantity
			expirationDTM := request.ExpirationDTM
			snackbatch := models.SnackBatch{
				SnackID: snackId,
				Quantity: quantity,
				ExpirationDTM: expirationDTM,
			}

			if err := tx.Model(models.SnackBatch{}).Create(&snackbatch).Error; err != nil {
				tx.Rollback()
				return err
			}
		}

		return nil
	})

	return snack, err
}

func (SnackService) UpdateSnack(request validators.SnackUpdateRequest, id int) (models.Snack, error) {
	db := database.GetDB()

	snack := models.Snack{}
	if err := db.First(&snack, id).Error; err != nil {
		return snack, err
	}

	if request == (validators.SnackUpdateRequest{}) {
		return snack, nil
	}

	// from request to snack model
	validators.UpdateRequestToSnackModel(request, &snack)

	if err := db.Save(&snack).Error; err != nil {
		return snack, err
	}

	return snack, nil
}

func (SnackService) DeleteSnack(id int) error {
	db := database.GetDB()

	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&models.Snack{}, id).Error; err != nil {
			return err
		}

		if err := tx.Where("snack_id = ?", id).Delete(&models.SnackBatch{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Delete(&models.Snack{}, id).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return err
}
