package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Product struct {
	id	int	`json: "id"`
    	stock   int	`json: "stock"`
    	name    string	`json: "name"`
    	color   string	`json: "color"`
    	code    string	`json: "code"`
    	price   float64 `json: "price"`
    	publish bool	`json: "publish"`
    	date    string	`json: "date"`
}


func main() {
	router := gin.Default()

	router.GET("/nombre", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello 👋 Bootcampers",
		})
	})

	router.GET("/products", GetAll)
	router.Run(":8000") // Por defecto gin arranca el server on 8080
}

func GetAll(ctx *gin.Context) {

	products := []Product{
		{id: 1, stock: 12, name: "shirt", color: "blue", code: "AC123", price: 12.4, publish: true, date: "December"},
		{id: 2, stock: 123, name: "pants", color: "red", code: "AC124", price: 3.4, publish: false, date: "November"},
	}

	ctx.JSON(http.StatusOK, gin.H{"message": products})
}

