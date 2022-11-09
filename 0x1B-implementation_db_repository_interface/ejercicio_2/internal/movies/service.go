package movies

import (
        "context"
        "github.com/bootcamp-go/storage/internal/domain"
)

type Service interface {
	Store(domain.Movie) (int, error)
	GetByName(name string) (domain.Movie, error)
        GetAll(c context.Context) ([]domain.Movie, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Store(p domain.Movie) (int, error) {
	return s.repository.Store(p)
}

func (s *service) GetAll(c context.Context) ([]domain.Movie, error) {
	movies, err := s.repository.GetAll(c)
	if err != nil {
		return []domain.Movie{}, err
	}

	return movies, err
}

func (s *service) GetByName(name string) (domain.Movie, error) {
	return s.repository.GetByName(name)
}
