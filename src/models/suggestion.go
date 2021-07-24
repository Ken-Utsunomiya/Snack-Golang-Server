package models

import (
	"time"
)

type Suggestion struct {
	ID 						 uint 		 `gorm:"column:suggestion_id;not null"`
	SuggestedBy 	 string 	 `gorm:"column:suggested_by;not null"`
	SuggestionText string 	 `gorm:"column:suggestion_text;not null"`
	SuggestionDTM  time.Time `gorm:"column:suggestion_dtm;not null"`
}

func (Suggestion) TableName() string {
	return "suggestions"
}
