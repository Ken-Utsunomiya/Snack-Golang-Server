package models

import (
	"time"
)

type Suggestion struct {
	ID		int		`json:"suggestion_id"`
	SuggestedBy 	string 	 	`json:"suggested_by"`
	SuggestionText 	string 	 	`json:"suggestion_text"`
	SuggestionDTM  	time.Time 	`json:"suggestion_dtm"`
}

func (Suggestion) TableName() string {
	return "suggestions"
}
