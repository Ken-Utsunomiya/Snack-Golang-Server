package services

import "Snack-Golang-Server/src/models"

type SnackService struct {}

func (SnackService) GetSnackList() []models.Snack {
	return nil
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
