package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
)

type SnackService struct {}

func (SnackService) GetSnackList() ([]models.Snack, error) {
	db := database.GetDB()
	snacks := make([]models.Snack, 0)
	err := db.Find(&snacks).Error
	return snacks, err
}

func (SnackService) AddSnack(snack *models.Snack) error {
	return nil
}

func (SnackService) UpdateSnack(snack *models.Snack) error {
	return nil
}

func (SnackService) DeleteSnack(id uint) error {
	return nil
}
