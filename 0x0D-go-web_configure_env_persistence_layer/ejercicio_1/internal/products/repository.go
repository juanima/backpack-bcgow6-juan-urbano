package products

import (
	"fmt"
        "errors"
	"github.com/ejercicio_1/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	GetOne(id int) (domain.Product, error)
	Store(name, color, code string, stock int, price float64, publish *bool) (domain.Product, error)
	FilterProduct(name, color, code string, id, stock int, price float64, publish *bool) ([]*domain.Product, error)
	Update(name, color, code string, id, stock int, price float64, publish *bool) (domain.Product, error)
	UpdateStock(id int, stock int) (domain.Product, error)
	Delete(id int) error
}

type repository struct{}


func NewRepository() Repository {
	return &repository{}
}

var (
	products []domain.Product
	lastID   int = 0
)

// Get one product
func (r *repository) GetOne(id int) (domain.Product, error) {

        var product domain.Product
        find := false

        for _, p := range products {
                if p.Id == id {
                        find = true
                        product = p
                        break
                 }
        }
      
        if !find {
                return product, errors.New("Not Found Product") 
        }

        return product, nil
}


// Get products
func (r *repository) GetAll() ([]domain.Product, error) {
	return products, nil
}

// Store de products
func (r *repository) Store(name, color, code string, stock int, price float64, publish *bool) (domain.Product, error) {
	p := domain.Product{
		Name: name,
		Color: color,
		Code: color,
		Stock: stock,
		Price: price,
		Publish: publish,
	}

	lastID++
	p.Id = lastID
	products = append(products, p)
	return p, nil
}

// Filter product
func (s *repository) FilterProduct(name, color, code string, id, stock int, price float64, publish *bool) ([]*domain.Product, error) {

        var filtrado []*domain.Product
        for _, t := range products { // filtrado por todos los campos
                if id == t.Id &&
                   stock == t.Stock &&
                   name == t.Name &&
                   color== t.Color &&
                   code == t.Code &&
                   price == t.Price &&
                   *publish == *t.Publish { 
                        filtrado = append(filtrado, &t)
                   }
        }
        
	return filtrado, nil 
}


// Update product
func (r *repository) Update(name, color, code string, id, stock int, price float64, publish *bool) (domain.Product, error) {
	p := domain.Product{
		Name: name,
		Color: color,
		Code: color,
		Stock: stock,
		Price: price,
		Publish: publish,
	}
	var updated bool
	for i := range products {
		if products[i].Id == id {
			p.Id = products[i].Id
			products[i] = p
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("product id %d not exists", id)
	}
	return p, nil
}



// Update Name product
func (r *repository) UpdateStock(id int, stock int) (domain.Product, error) {
	var updated bool
	p := domain.Product{}
	for i := range products {
		if products[i].Id == id {
			products[i].Stock = stock
			p = products[i]
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("product id %d not exists", id)
	}
	return p, nil
}


// Delete product
func (r *repository) Delete(id int) error {
	var deleted bool
	var pos int
	for i := range products {
		if products[i].Id == id {
			pos = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("product id %d not exists", id)
	}

	products = append(products[:pos], products[pos+1:]...)
	return nil
}

