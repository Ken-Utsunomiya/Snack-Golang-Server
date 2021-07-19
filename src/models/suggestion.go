package models

import "github.com/golang/protobuf/ptypes/timestamp"

type Suggestion struct {
	SuggestionId int `json:"suggestionId"`
	SuggestedBy string `json:"suggestedBy"`
	SuggestionText string `json:"suggestionText"`
	SuggestionDTM timestamp.Timestamp `json:"suggestionDTM"`
}
