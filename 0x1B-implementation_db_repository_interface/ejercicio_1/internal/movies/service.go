package movies

import "github.com/bootcamp-go/storage/internal/domain"

type Service interface {
	Store(domain.Movie) (int, error)
	GetByName(name string) (domain.Movie, error)
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

func (s *service) GetByName(name string) (domain.Movie, error) {
	return s.repository.GetByName(name)
}
