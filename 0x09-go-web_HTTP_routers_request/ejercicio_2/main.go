package main

import (
	"strconv"
	"net/http"
    	"github.com/gin-gonic/gin"
)


type Product struct {
	ID      int	`form:"id"`
	Stock   int	`form:"stock"`
    	Name    string  `form:"name"`
    	Color   string  `form:"color"`
    	Code    string  `form:"code"`
    	Price   float64 `form:"price"`
    	Publish bool    `form:"publish"`
    	Date    string  `form:"date"`
}

var products []Product


func ProductHandler(c *gin.Context) {
	var product Product 
	
	listProducts := generarProducts()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	find := false
	for _, p := range listProducts {
		if p.ID == int(id) {
			find = true
			product = p
			break
		}
	}

	if !find {
		c.JSON(http.StatusNotFound, nil) // 404
	} else {
		c.JSON(http.StatusOK, product) // 200 => transaction
	}
}



func generarProducts() []Product {
	products := []Product{
		{ID: 1, Stock: 12, Name: "shirt", Color: "blue", Code: "AC123", Price: 12.4, Publish: true, Date: "December"},
		{ID: 2, Stock: 123, Name: "pants", Color: "red", Code: "AC124", Price: 3.4, Publish: false, Date: "November"},
	}
	return products
}


func main() {

	router := gin.Default()
	router.GET("/products/:id", ProductHandler)

	router.Run(":8000") // Por defecto gin arranca el server on 8080
}
