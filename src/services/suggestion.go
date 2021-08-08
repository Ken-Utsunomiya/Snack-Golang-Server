package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
	"gorm.io/gorm"
)

type SuggestionService struct {}

func (SuggestionService) GetSuggestionList() ([]models.Suggestion, error) {
	db := database.GetDB()
	suggestions := make([]models.Suggestion, 0)
	err := db.Find(&suggestions).Error
	return suggestions, err
}

func (SuggestionService) AddSuggestion(suggestion *models.Suggestion) error {
	return nil
}

func (SuggestionService) DeleteSuggestion() error {
	db := database.GetDB()
	err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Suggestion{}).Error
	return err
}
