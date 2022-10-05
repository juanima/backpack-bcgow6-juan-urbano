package main

import (
	"net/http"
    	"github.com/gin-gonic/gin"
)


type request struct {
	ID      int	`json: "id"`
	Stock   int	`json: "stock"`
    	Name    string  `json: "name"`
    	Color   string  `json: "color"`
    	Code    string  `json: "code"`
    	Price   float64 `json: "price"`
    	Publish bool    `json: "publish"`
    	Date    string  `json: "date"`
}

var products []request


func saveProduct() gin.HandlerFunc {
	return func (c *gin.Context) {
	
		var req request

		token := c.Request.Header.Get("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

    	    	if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	    	    	return
    	    	}

		if req.Stock == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El numero stock es requerido"})
			return
		}
		if req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El name es requerido"})
			return
		}
		if req.Color == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El color requerida"})
			return
		}
		if req.Code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El código es requerido"})
			return
		}
		if req.Price == 0.0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El precio es requerido"})
			return
		}
		if req.Publish == false {
			c.JSON(http.StatusBadRequest, gin.H{"error": "El publish es requerido"})
			return
		}
		if req.Date == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "La fecha es requerido"})
			return
		}

    	    	req.ID = len(products) + 1
    	    	products = append(products, req)
    	    	c.JSON(http.StatusCreated, req)
    	}
}


func main() {
	r := gin.Default()
    	pr := r.Group("/products")
    	pr.POST("/", saveProduct())
    	r.Run()
}
