package handlers

import (
	"Snack-Golang-Server/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuggestionList(c *gin.Context) {
	suggestionService := services.SuggestionService{}
	suggestionList, err := suggestionService.GetSuggestionList()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{ "suggestions": suggestionList })
}

func SuggestionCreate(c *gin.Context) {

}

func SuggestionDelete(c *gin.Context) {

}
