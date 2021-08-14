package validators

import (
	"Snack-Golang-Server/src/models"
	"time"
)

type SuggestionRegisterRequest struct {
	SuggestedBy int `json:"suggested_by" binding:"required"`
	SuggestionText string `json:"suggestion_text" binding:"required"`
}

func RegisterRequestToSuggestionModel(request SuggestionRegisterRequest) models.Suggestion {
	suggestion := models.Suggestion{}
	suggestion.SuggestedBy = request.SuggestedBy
	suggestion.SuggestionText = request.SuggestionText
	suggestion.SuggestionDTM = time.Now()
	return suggestion
}
