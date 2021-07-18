package main

import (
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

		routers.UserRoutes(v1)
		routers.TransactionRoutes(v1)
	}
	r.router.Run(":3000")
}
