package handlers

import (
	"Snack-Golang-Server/src/middlewares"
	"Snack-Golang-Server/src/services"
	"Snack-Golang-Server/src/validators"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuggestionList(c *gin.Context) {
	suggestionService := services.SuggestionService{}
	suggestionList, err := suggestionService.GetSuggestionList()
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{ "suggestions": suggestionList })
}

func SuggestionCreate(c *gin.Context) {
	suggestionService := services.SuggestionService{}

	registerRequest := validators.SuggestionRegisterRequest{}

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		_ = c.Error(errors.New(middlewares.BadRequest)).SetType(gin.ErrorTypePublic)
		return
	}

	res, err := suggestionService.AddSuggestion(registerRequest)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func SuggestionDelete(c *gin.Context) {
	suggestionService := services.SuggestionService{}
	if err := suggestionService.DeleteSuggestion(); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.Status(http.StatusNoContent)
}
