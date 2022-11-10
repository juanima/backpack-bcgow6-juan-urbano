package main

import (
	"github.com/bootcamp-go/storage/cmd/server/handler"
	cnn "github.com/bootcamp-go/storage/db"
	"github.com/bootcamp-go/storage/internal/movies"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db := cnn.MySQLConnection()
	repo := movies.NewRepository(db)
	serv := movies.NewService(repo)
	p := handler.NewMovie(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/movies")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetByName())
	pr.GET("", p.GetAll())
        pr.DELETE("/:id", p.Delete())

	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
