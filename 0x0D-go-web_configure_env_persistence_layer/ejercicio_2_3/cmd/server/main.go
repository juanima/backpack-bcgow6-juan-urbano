package main

import (
        "log"
	"github.com/gin-gonic/gin"
        "github.com/joho/godotenv"
	"github.com/ejercicio_2/cmd/server/handler"
	"github.com/ejercicio_2/internal/products"
        "github.com/ejercicio_2/pkg/store"
)


func main() {

        err := godotenv.Load()
        if err != nil {
          log.Fatal("error al intentar cargar archivo .env")
        }

        db := store.New(store.FileType, "./products.json")

	repo := products.NewRepository(db)
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
