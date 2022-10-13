package handler

import (
        "os"
        "log"
        "errors"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/ejercicio_1/internal/products"
)

const (
        TOKEN = "TOKEN"
)

type ProductRequest struct {
        Id      int     `form:"id"      json:"id"`
	Stock   int     `form:"stock"   json:"stock"   binding:"required"`
	Name    string  `form:"name"    json:"name"    binding:"required"`
	Color   string  `form:"color"   json:"color"   binding:"required"`
	Code    string  `form:"code"    json:"code"    binding:"required"`
	Price   float64 `form:"price"   json:"price"   binding:"required"`
	Publish *bool   `form:"publish" json:"publish" binding:"required"`
	Date    string  `form:"date"    json:"date"    binding:"required"`
}


type ProductRequestPatch struct {
	Stock int `json:"stock" binding:"required"`
}

type Product struct {
	service products.Service
}

func NewProduct(s products.Service) *Product {
	return &Product{service: s}
}

func getToken(name string) (string, error) {

        token := os.Getenv(TOKEN) 
        if token == "" {
                return token, errors.New("Token no encontrado")
        }
        
        return token, nil
}

func (p *Product) Status() gin.HandlerFunc {
	return func(c *gin.Context) {
                c.JSON(http.StatusOK, gin.H{
                        "message": "it works!!",
                })
        }
}


func (p *Product) GetOne() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return
		}

                product, err := p.service.GetOne(id)
                if err != nil {
                        c.JSON(http.StatusNotFound, nil) 
                        return
                }

                c.JSON(http.StatusOK, product) 
	}
}


func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
                
                token_system, err := getToken(token)
                if err != nil || token != token_system {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
                        return
                }

		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(p) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "no hay products registrados"})
			return
		}

		c.JSON(http.StatusOK, p)
		// _ = c
	}
}


func (p *Product) FilterProduct() gin.HandlerFunc {
        return func(c *gin.Context) {

                var req ProductRequest

                if c.ShouldBindQuery(&req) == nil { // Setea las variables obtenidas de c.Query("nombredelavariable")
                        log.Println(req.Id)
                        log.Println(req.Stock)
                        log.Println(req.Name)
                        log.Println(req.Color)
                        log.Println(req.Code)
                        log.Println(req.Price)
                        log.Println(req.Publish)
                }

		filtrado, err := p.service.FilterProduct(req.Name, req.Color, req.Code, req.Id, req.Stock, req.Price, req.Publish)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

                c.JSON(http.StatusOK, filtrado) // devovemos el array filtrado
        }
}


func (p *Product) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ProductRequest
	
                token := c.GetHeader("token")

                token_system, err := getToken(token)
                if err != nil || token != token_system {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
                        return
                }
                

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo nombre es requerido"})
			return
		}

		if req.Color == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo color es requerido"})
			return
		}

		if req.Code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo código es requerido"})
			return
		}

		if req.Stock <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo stock es requerido"})
			return
		}

		if req.Price <= 0.0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo precio es requerido"})
			return
		}

                *req.Publish = false

		p, err := p.service.Store(req.Name, req.Color, req.Code, req.Stock, req.Price, req.Publish)
		if err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, p)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

                token_system, err := getToken(token)
                if err != nil || token != token_system {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
                        return
                }

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id invalido - " + err.Error()})
			return
		}

		var req ProductRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Name == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo nombre es requerido"})
			return
		}

		if req.Color == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo color es requerido"})
			return
		}

		if req.Code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo codigo es requerido"})
			return
		}

		if req.Date == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo fecha es requerido"})
			return
		}

		if req.Stock <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo stock es requerido"})
			return
		}

		if req.Price <= 0.0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo precio es requerido"})
			return
		}

		if req.Publish == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "campo publicado es requerido"})
			return
		}

		p, err := p.service.Update(req.Name, req.Color, req.Code, id, req.Stock, req.Price, req.Publish)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, p)
	}
}

func (p *Product) UpdateStock() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
                
                token_system, err := getToken(token)
                if err != nil || token != token_system {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
                        return
                }

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id invalido - " + err.Error()})
			return
		}

		var req ProductRequestPatch
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		p, err := p.service.UpdateStock(id, req.Stock)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, p)
        }
}


func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

	        token := c.GetHeader("token")


                token_system, err := getToken(token)
                if err != nil || token != token_system {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
                        return
                }
	        // if token != "123456" || token == "" {
	        // 	c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
	        // 	return
	        // }

	        id, err := strconv.Atoi(c.Param("id"))
	        if err != nil {
	        	c.JSON(http.StatusBadRequest, gin.H{"error": "Id invalido - " + err.Error()})
	        	return
	        }

	        if err := p.service.Delete(id); err != nil {
	        	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	        	return
	        }

	        c.JSON(http.StatusNoContent, gin.H{"message": "algo"})
        }
}
