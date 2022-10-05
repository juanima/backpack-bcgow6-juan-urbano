package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/nombre", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hola 👋 Bootcampers",
		})
	})

	router.Run(":8000") // Por defecto gin arranca el server on 8080
}
