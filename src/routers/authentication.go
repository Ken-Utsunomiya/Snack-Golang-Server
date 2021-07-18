package routers

import (
	"Snack-Golang-Server/src/handlers"
	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(rg *gin.RouterGroup) {

	authentication := rg.Group("/authenticate")
	{
		authentication.POST("/", handlers.VerifyAndCreateToken)
	}
}