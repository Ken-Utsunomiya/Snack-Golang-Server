package routers

import (
	"Snack-Golang-Server/src/handlers"
	"github.com/gin-gonic/gin"
)

func SuggestionRoutes(rg *gin.RouterGroup) {

	suggestion := rg.Group("/suggestions")
	{
		suggestion.GET("/", handlers.SuggestionList)
		suggestion.POST("/", handlers.SuggestionCreate)
		suggestion.DELETE("/", handlers.SuggestionDelete)
	}
}
