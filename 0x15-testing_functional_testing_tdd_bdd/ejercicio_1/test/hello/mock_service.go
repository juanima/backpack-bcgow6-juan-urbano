package hello

import (
	"fmt"
	"github.com/ejercicio_2/internal/domain"
)

type MockService struct {
	DataMock []domain.Product
	Error string
}

func (m *MockService) GetAll() ([]domain.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}

	return m.DataMock, nil
}

func (m *MockService) GetOne(id int) (domain.Product, error) {
        return domain.Product{}, nil
}


func (m *MockService) FilterProduct(name, color, code string, id, stock int, price float64, publish *bool) ([]*domain.Product, error) {
        return nil, nil
}

func (m *MockService) Store(name, color, code string, stock int, price float64, publish *bool) (domain.Product, error) {
        if m.Error != "" {
		return domain.Product{}, fmt.Errorf(m.Error)
	}

        var lastID int
	p := domain.Product{
		Name: name,
		Color: color,
		Code: code,
		Stock: stock,
		Price: price,
		Publish: publish,
	}

	if len(m.DataMock) == 0 {
		lastID = 0 
	} else {
	        lastID = m.DataMock[len(m.DataMock) - 1].Id
        }

	lastID++
	p.Id = lastID
	m.DataMock = append(m.DataMock, p)

	return p, nil
}



func (m *MockService) Update(name, color, code string, id, stock int, price float64, publish *bool) (domain.Product, error) {

        product := domain.Product{
		Name: name,
		Color: color,
		Code: code,
		Stock: stock,
		Price: price,
		Publish: publish,
	}

	var updated bool
	for i := range m.DataMock {
		if m.DataMock[i].Id == id {
			product.Id = m.DataMock[i].Id
			m.DataMock[i] = product
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("product id %d not exists", id)
	}

        return product, nil
}

func (m *MockService) UpdateStock(id int, stock int) (domain.Product, error){
        return domain.Product{}, nil
}

func (m *MockService) UpdateName(id int, name string) (domain.Product, error){
        return domain.Product{}, nil
}

func (m *MockService) Delete(id int) error  {
        return nil
}
