package main

import (
	"fmt"
	"log"
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


func FiltrarProductsHandler(c *gin.Context) {

	listaProducts := generarProducts()
	var product Product

	if c.ShouldBindQuery(&product) == nil { // Setea las variables obtenidas de c.Query("nombredelavariable")
		log.Println(product.ID)
		log.Println(product.Stock)
		log.Println(product.Name)
		log.Println(product.Color)
		log.Println(product.Code)
		log.Println(product.Price)
		log.Println(product.Publish)
		log.Println(product.Date)
	}

	fmt.Printf("%+v\n", product)

	var filtrado []*Product
	for _, t := range *listaProducts { // filtrado por todos los campos
		if product.ID == t.ID && 
		   product.Stock == t.Stock && 
		   product.Name == t.Name && 
		   product.Color== t.Color && 
		   product.Code == t.Code && 
		   product.Price == t.Price && 
		   product.Publish == t.Publish &&
		   product.Date == t.Date {
			filtrado = append(filtrado, &t)
		}
	}

	fmt.Printf("%+v\n", filtrado)

	c.JSON(http.StatusOK, filtrado) // devovemos el array filtrado
}

func generarProducts() *[]Product {
	products := []Product{
		{ID: 1, Stock: 12, Name: "shirt", Color: "blue", Code: "AC123", Price: 12.4, Publish: true, Date: "December"},
		{ID: 2, Stock: 123, Name: "pants", Color: "red", Code: "AC124", Price: 3.4, Publish: false, Date: "November"},
	}
	return &products
}



func main() {
	router := gin.Default()
	router.GET("/filter-product", FiltrarProductsHandler)

	router.Run(":8000") // Por defecto gin arranca el server on 8080
}
