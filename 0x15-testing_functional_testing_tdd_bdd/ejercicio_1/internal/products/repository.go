package products

import (
	"fmt"
        "errors"
	"github.com/ejercicio_2/internal/domain"
        "github.com/ejercicio_2/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	GetOne(id int) (domain.Product, error)
        Store(name, color, code string, stock int, price float64, publish *bool) (domain.Product, error)
	FilterProduct(name, color, code string, id, stock int, price float64, publish *bool) ([]*domain.Product, error)
	Update(name, color, code string, id, stock int, price float64, publish *bool) (domain.Product, error)
	UpdateStock(id int, stock int) (domain.Product, error)
	UpdateName(id int, name string) (domain.Product, error)
	Delete(id int) error
}

type repository struct{
        db store.Store
}


func NewRepository(db store.Store) Repository {
	return &repository{
                db: db,
        }
}

var (
	// products []domain.Product
	// lastID   int = 0
)

 

// Get one product
func (r *repository) GetOne(id int) (product domain.Product, err error) {

        var products []domain.Product
        find := false
        
        err = r.db.Read(&products)
        if err != nil {
                return product, err
        }

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
func (r *repository) GetAll() (products []domain.Product, err error) {

        err = r.db.Read(&products)
        if err != nil {
                return nil, err
        }

	return products, nil
}

// Store de products
func (r *repository) Store(name, color, code string, stock int, price float64, publish *bool) (product domain.Product, err error) {

        var products []domain.Product
        var lastID int

        err = r.db.Read(&products)
        if err != nil {
                return product, err
        }

	p := domain.Product{
		Name: name,
		Color: color,
		Code: code,
		Stock: stock,
		Price: price,
		Publish: publish,
	}

	if len(products) == 0 {
		lastID = 0 
	} else {
	        lastID = products[len(products) - 1].Id
        }

	lastID++
	p.Id = lastID
	products = append(products, p)

        err = r.db.Write(products)
        if err != nil {
                return product, err
        }

	return p, nil
}

// Filter product
func (r *repository) FilterProduct(name, color, code string, id, stock int, price float64, publish *bool) (products []*domain.Product, err error) {

        var filtrado []*domain.Product

        err = r.db.Read(&products)
        if err != nil {
                return products, err
        }

        for _, t := range products { // filtrado por todos los campos
                if id == t.Id &&
                   stock == t.Stock &&
                   name == t.Name &&
                   color== t.Color &&
                   code == t.Code &&
                   price == t.Price &&
                   *publish == *t.Publish { 
                        filtrado = append(filtrado, t)
                   }
        }
        
	return filtrado, nil 
}


// Update product
func (r *repository) Update(name, color, code string, id, stock int, price float64, publish *bool) (product domain.Product, err error) {

        var products []domain.Product

        err = r.db.Read(&products)
        if err != nil {
                return product, err
        }

	product = domain.Product{
		Name: name,
		Color: color,
		Code: code,
		Stock: stock,
		Price: price,
		Publish: publish,
	}

	var updated bool
	for i := range products {
		if products[i].Id == id {
			product.Id = products[i].Id
			products[i] = product
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("product id %d not exists", id)
	}
        
        err = r.db.Write(products)
        // err = r.db.Write(product)
        if err != nil {
                return product, err
        }

	return product, nil
}



// Update Name product
func (r *repository) UpdateStock(id int, stock int) (product domain.Product, err error) {
	var updated bool
        var products []domain.Product

        err = r.db.Read(&products)
        if err != nil {
                return product, err
        }

	product = domain.Product{}
	for i := range products {
		if products[i].Id == id {
			products[i].Stock = stock
			product = products[i]
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("product id %d not exists", id)
	}

        err = r.db.Write(products)
        if err != nil {
                return product, err
        }

        return product, nil
}


// Update Name product
func (r *repository) UpdateName(id int, name string) (product domain.Product, err error) {
	var updated bool
        var products []domain.Product

        err = r.db.Read(&products)
        if err != nil {
                return product, err
        }

	product = domain.Product{}
	for i := range products {
		if products[i].Id == id {
			products[i].Name = name 
			product = products[i]
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("product id %d not exists", id)
	}

        err = r.db.Write(products)
        if err != nil {
                return product, err
        }

        return product, nil
}


// Delete product
func (r *repository) Delete(id int) error {
	var deleted bool
	var pos int

        var products []domain.Product

        err := r.db.Read(&products)
        if err != nil {
                return err
        }

	for i := range products {
		if products[i].Id == id {
			pos = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("product id %d not exists", id)
	}

        products = append(products[:pos], products[pos + 1:]...)

        err = r.db.Write(products)
        if err != nil {
                return err
        }
        
	return nil
}

