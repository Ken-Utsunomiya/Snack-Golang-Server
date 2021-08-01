package models

import (
	"time"
)

type Suggestion struct {
	ID 						 int 		   `json:"suggestion_id"   gorm:"column:suggestion_id;primaryKey;not null"`
	SuggestedBy 	 string 	 `json:"suggested_by"    gorm:"column:suggested_by;not null"`
	SuggestionText string 	 `json:"suggestion_text" gorm:"column:suggestion_text;not null"`
	SuggestionDTM  time.Time `json:"suggestion_dtm"  gorm:"column:suggestion_dtm;not null"`
}

func (Suggestion) TableName() string {
	return "suggestions"
}
