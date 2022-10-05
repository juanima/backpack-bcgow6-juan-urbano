package main

import (
	"net/http"
    	"github.com/gin-gonic/gin"
)

/*
if err := ctx.ShouldBindJSON(&userRequest); err != nil { //GET ALL THE STRUCT USER BY BODY
	if userRequest.Name == "" || userRequest.Surname == "" || userRequest.Email == "" || userRequest.Age == 0 || userRequest.Height == 0 || userRequest.CreateDate.IsZero() {
	if userRequest.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Field name is required!"})
	}
	if userRequest.Surname == "" {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Field surname is required!"})
	}
*/


type request struct {
	ID      int	`json: "id"`
	stock   int	`json: "stock"`
    	Name    string  `json: "name"`
    	Color   string  `json: "color"`
    	Code    string  `json: "code"`
    	Price   int	`json: "price"`
    	Publish bool    `json: "publish"`
    	Date    string  `json: "date"`
}

var products []request

func saveProduct() gin.HandlerFunc {
	return func (c *gin.Context) {
		var req request
    	    	if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(404, gin.H{
				"error": err.Error(),
    	    	    	})
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
