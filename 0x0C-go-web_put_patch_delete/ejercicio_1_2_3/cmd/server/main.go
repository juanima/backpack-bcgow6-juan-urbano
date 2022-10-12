package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ejercicio_1/cmd/server/handler"
	"github.com/ejercicio_1/internal/products"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
        r.GET("/status", p.Status())

	pr := r.Group("products")
	pr.GET("/", p.GetAll())
        pr.GET("/:id", p.GetOne())
        pr.GET("/filter-product", p.FilterProduct())       
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateStock())
	pr.DELETE("/:id", p.Delete())

	if err := r.Run(); err != nil {
		panic(err)
	}
}
