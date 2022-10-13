package products

import "github.com/ejercicio_1/internal/domain"

type Service interface {
	GetAll() ([]domain.Product, error)
	GetOne(id int) (domain.Product, error)
	FilterProduct(name, color, code string, id, stock int, price float64, publish *bool) ([]*domain.Product, error)
	Store(name, color, code string, stock int, price float64, publish *bool) (domain.Product, error)
	Update(name, color, code string, id, stock int, price float64, publish *bool) (domain.Product, error)
	UpdateStock(id int, stock int) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetOne(id int) (domain.Product, error) {
        return s.repository.GetOne(id)
}

func (s *service) GetAll() ([]domain.Product, error) {
	return s.repository.GetAll()
}

func (s *service) FilterProduct(name, color, code string, id, stock int, price float64, publish *bool) ([]*domain.Product, error) {
	return s.repository.FilterProduct(name, color, code, id, stock, price, publish) 
}

func (s *service) Store(name, color, code string, stock int, price float64, publish *bool) (domain.Product, error) {
	return s.repository.Store(name, color, code, stock, price, publish)
}

func (s *service) Update(name, color, code string, id, stock int, price float64, publish *bool) (domain.Product, error) {
	return s.repository.Update(name, color, code, id, stock, price, publish)
}

func (s *service) UpdateStock(id int, stock int) (domain.Product, error) {
	return s.repository.UpdateStock(id, stock)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
