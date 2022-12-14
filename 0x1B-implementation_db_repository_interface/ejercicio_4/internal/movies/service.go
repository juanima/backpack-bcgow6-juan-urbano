package movies

import (
        "context"
        "github.com/bootcamp-go/storage/internal/domain"
)

type Service interface {
	Store(context.Context, domain.Movie) (int, error)
	GetByName(name string) (domain.Movie, error)
        GetAll(c context.Context) ([]domain.Movie, error)
        Delete(c context.Context, id int64) error
        Update(ctx context.Context, b domain.Movie, id int) (domain.Movie, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Store(c context.Context, p domain.Movie) (int, error) {
	return s.repository.Store(c, p)
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

func (s *service) Delete(c context.Context, id int64) error {
	return s.repository.Delete(c, id)
}

func (s *service) Update(ctx context.Context, b domain.Movie, id int) (domain.Movie, error) {

	err := s.repository.Update(ctx, b, id)
	if err != nil {
		return domain.Movie{}, err
	}
	updated, err := s.repository.GetMovieByID(ctx, id)
	if err != nil {
		return b, err
	}
	return updated, nil
}

