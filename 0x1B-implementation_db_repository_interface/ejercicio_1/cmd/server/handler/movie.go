package handler

import (
        "time"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"

	"github.com/bootcamp-go/storage/internal/domain"
	"github.com/bootcamp-go/storage/internal/movies"
	"github.com/bootcamp-go/storage/pkg/web"
)


type Request struct {
	ID           int       `json:"id"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
	Title        string    `json:"title"`
	Rating       float32   `json:"rating"`
	Awards       int       `json:"awards"`
	Release_date string    `json:"release_date"`
	Length       *int      `json:"length"`
	Genre_id     *int      `json:"genre_id"`
}


type requestName struct {
	Title string `json:"title" binding:"required"`
}

type Movie struct {
	service movies.Service
}

func NewMovie(s movies.Service) *Movie {
	return &Movie{service: s}
}

func (s *Movie) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		product := domain.Movie(req)
		id, err := s.service.Store(product)
		if err != nil {
			c.JSON(http.StatusConflict, err.Error())
			return
		}

		product.ID = id
		c.JSON(http.StatusCreated, web.NewResponse(product, "", http.StatusCreated))
	}
}

func (s *Movie) GetByName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestName
		if err := c.ShouldBindJSON(&req); err != nil {
			if strings.Contains(err.Error(), "'required' tag") {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			c.JSON(http.StatusBadRequest, web.NewResponse(nil, err.Error(), http.StatusBadRequest))
			return
		}

		product, err := s.service.GetByName(req.Title)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(nil, err.Error(), http.StatusNotFound))
			return
		}

		c.JSON(http.StatusOK, web.NewResponse(product, "", http.StatusOK))
	}
}


