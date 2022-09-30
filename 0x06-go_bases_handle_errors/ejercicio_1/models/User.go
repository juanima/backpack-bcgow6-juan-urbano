package models 

import "fmt"

type User struct {
    name     string	`json:"ID"`
    lastName string	`json:"price"`
    email    string	`json:"email"`
    products []*Product `json:"password`
}


func NewUser(name, lastName, email string) (user *User, err error) {

    // products := [...]*Product{} 

    user = &User {
	name: name, 
	lastName: lastName,
	email: email, 
	products: nil, //products, 
    }

    return user, nil
}

func (u *User) StrUser() string {
    infoUser := fmt.Sprintf(
	"Info User: \n\tName: %s,\n\tLast name: %s, \n\tEmail: %s, \n\tProducts: \n", 
		    u.name, u.lastName, u.email)   

    for _, product := range u.products {
	// infoUser += (*product).StrProduct()
	fmt.Println(product.StrProduct())
    }

    return infoUser
}


func (u *User) SetName(newName string) {
    u.name = newName
}


func (u *User) SetEmail(newEmail string) {
    u.email = newEmail
}


func (u *User) SetLastName(newLastName string) {
    u.lastName = newLastName
}


func (u *User) SetProducts(newProducts *[]*Product) {
    u.products = *newProducts 
}


func (u *User) GetProducts() []*Product {
    return u.products
}

func (u *User) GetName() string {
    return u.name 
}


func (u *User) GetLastName() string {
    return u.lastName
}


