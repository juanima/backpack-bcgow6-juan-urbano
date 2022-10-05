package service

type Carts interface {
	// Create create a new Product
	Create(p Product) (Product, error)
}

type carts struct {
	Products []Product
}

type Product struct {
    Id	    int		`json: "id"`
    stock   int		`json: "stock"`
    Name    string	`json: "name"`
    Color   string	`json: "color"`
    Code    string	`json: "code"`
    Price   int		`json: "price"`
    Publish bool	`json: "publish"`
    Date    string	`json: "date"`
}

// NewBookings creates a new carts service
func NewCarts(products []Product) Carts {
	return &carts{Products: products }
}

func (b *carts) Create(t Product) (Product, error) {
	return Product{}, nil
}

