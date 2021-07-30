package services

import "Snack-Golang-Server/src/models"

type SuggestionService struct {}

func (SuggestionService) GetSuggestionList() []models.Suggestion {
	return nil
}

func (SuggestionService) AddSuggestion(suggestion *models.Suggestion) error {
	return nil
}

func (SuggestionService) DeleteSuggestion(id uint) error {
	return nil
}
