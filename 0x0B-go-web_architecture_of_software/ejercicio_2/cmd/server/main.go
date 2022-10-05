package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ejercicio_2/cmd/server/handler"
	"github.com/ejercicio_2/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)

	p := handler.NewProduct(service)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	r.Run()
}
