package main

import (
        "log"
        "os"
	"github.com/gin-gonic/gin"
        "github.com/joho/godotenv"
	"github.com/ejercicio_2/cmd/server/handler"
	"github.com/ejercicio_2/internal/products"
        "github.com/ejercicio_2/pkg/store"
        "github.com/ejercicio_2/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @title           Bootcamp Go Wave 6 - API
// @version         1.0
// @description     This is a simple API development by Digital House's team.
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name   API Support Digital House
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      http://127.0.0.1:8080
// @BasePath  /api/v1
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
	api := r.Group("/api/v1")

        api.Use(handler.AuthMiddlewareToken())

        // Documentación swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

        api.GET("/status", p.Status())

	pr := api.Group("products") 
        {
	        pr.GET("/", p.GetAll())
                pr.GET("/:id", p.GetOne())
                pr.GET("/filter-product", p.FilterProduct())       
	        pr.POST("/", p.Store())
	        pr.PUT("/:id", p.Update())
	        pr.PATCH("/:id", p.UpdateStock())
	        pr.DELETE("/:id", p.Delete())

        }

	if err := r.Run(); err != nil {
		panic(err)
	}
}
