package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ejercicio_2/pkg/web"
)

func AuthMiddlewareToken() gin.HandlerFunc {
	tokenAPI := os.Getenv("TOKEN")
	if tokenAPI == "" {
		log.Fatal("No se ha registrado el token para la API")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if tokenAPI != token {
			c.AbortWithStatusJSON(http.StatusForbidden, web.NewResponse(http.StatusForbidden, nil, "Forbidden 403"))
			//c.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		c.Next()
	}
}
