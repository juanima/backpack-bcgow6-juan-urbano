package handler

import (
	"log"
	"os"
        "errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ejercicio_2/pkg/web"
)

const (
        TOKEN_ENV = "TOKEN"
        TOKEN_HEADER = "token"
)

func getToken(name string) (string, error) {

        token := os.Getenv(name) 
        if token == "" {
                return token, errors.New("Token no encontrado")
        }
        
        return token, nil
}

func AuthMiddlewareToken() gin.HandlerFunc {
	tokenAPI, err := getToken(TOKEN_ENV) 
        if err != nil {
		log.Fatal("No se ha registrado el token para la API")
        }

	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN_HEADER)
		if tokenAPI != token {
			c.AbortWithStatusJSON(http.StatusForbidden, web.NewResponse(http.StatusForbidden, nil, "Forbidden 403"))
			//c.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token inválido"))
			return
		}

		c.Next()
	}
}
