package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/middlewares"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/validators"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type SuggestionService struct {}

func (SuggestionService) GetSuggestionList() ([]models.Suggestion, error) {
	db := database.GetDB()
	suggestions := make([]models.Suggestion, 0)
	err := db.Find(&suggestions).Error
	return suggestions, err
}

func (SuggestionService) AddSuggestion(request validators.SuggestionRegisterRequest) (models.Suggestion, error) {
	db := database.GetDB()
	var suggestion models.Suggestion

	err := db.Transaction(func(tx *gorm.DB) error {
		userId := request.SuggestedBy
		if err := tx.Model(&models.User{}).First(&models.User{}, models.User{ID: userId}).Error; err != nil {
			return err
		}

		text := strings.TrimSpace(request.SuggestionText)
		if len(text) == 0 {
			return errors.New(middlewares.BadRequest)
		}
		request.SuggestionText = text
		suggestion = validators.RegisterRequestToSuggestionModel(request)

		if err := tx.Model(&models.Suggestion{}).Create(&suggestion).Error; err != nil {
			tx.Rollback()
			return err
		}

		return nil
	})

	return suggestion, err
}

func (SuggestionService) DeleteSuggestion() error {
	db := database.GetDB()
	err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Suggestion{}).Error
	return err
}
