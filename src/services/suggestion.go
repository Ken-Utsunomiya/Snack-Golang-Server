package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
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

func (SuggestionService) DeleteSuggestion(id uint) error {
	return nil
}
