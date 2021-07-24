package main

import (
	"Snack-Golang-Server/src/database"
	"github.com/gin-gonic/gin"

	"Snack-Golang-Server/src/routers"
)

type routes struct {
	router *gin.Engine
}

func main() {

	r := routes{
		router: gin.Default(),
	}

	apiEngine := r.router.Group("/api")
	{
		v1 := apiEngine.Group("/v1")

		routers.AuthenticationRoutes(v1)
		routers.UserRoutes(v1)
		routers.TransactionRoutes(v1)
		routers.SuggestionRoutes(v1)
		routers.PaymentRoutes(v1)
		routers.SnackRoutes(v1)
		routers.SnackBatchRoutes(v1)
	}

	db := database.Init()
	database.AutoMigrate(db)

	r.router.Run(":3000")
}
