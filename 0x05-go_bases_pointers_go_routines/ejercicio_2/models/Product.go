package models 

import "fmt"

type Product struct {
    name   string  `json:"name"`
    price  float64 `json:"price"`
    amount int     `json:"amount"`
}


func NewProduct(name string, price float64) (product *Product, err error) {

    product = &Product {
	name: name, 
	price: price,
	amount: 1,
    }

    return product, nil
}

func (p *Product) StrProduct() string {
    return fmt.Sprintf(
	"Info Product: \n\tName: %s,\n\tPrice: %f, \n\tAmount: %d\n", 
	p.name, p.price, p.amount)    
}


func (p *Product) SetName(newName string) {
    p.name = newName
}


func (p *Product) SetPrice(newPrice float64) {
    p.price = newPrice
}


func (p *Product) SetAmount(newAmount int) {
    if p.amount > 0 {
	p.amount = newAmount
    }
}

func (p *Product) GetAmount() int {
    return p.amount
}

func (p *Product) GetName() string {
    return p.name
}

func (p *Product) GetPrice() float64 {
    return p.price
}







