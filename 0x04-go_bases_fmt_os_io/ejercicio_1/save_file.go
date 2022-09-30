package main

import (
    "fmt"
    "os"
)

const (
    FILENAME = "./myFile.csv" 
    PERMISSION = 0644
)

type Product struct {
    Id int	    `json:"ID"`
    Price float64   `json:"price"`
    Amount int	    `json:"amount"`
}


func detail(products []Product, numberProducts int) string {

    var msg string
    var input string

    for i := 0; i < numberProducts; i++ {
	
	msg = fmt.Sprintf("%d,%.2f,%d\n", products[i].Id, products[i].Price, products[i].Amount)
	input += msg
    }

    // fmt.Printf("Info del producto %s \n", msg)
    return input
}


func serializeProduct(id, amount int, price float64) (product Product, err error) {

    product = Product {
	Id: id, 
	Price: price, 
	Amount: amount, 
    }

    return product, nil
}

func saveProducts(text string) {

   d1 := []byte(text)
   err := os.WriteFile(FILENAME, d1, PERMISSION)
   if err != nil {
	panic(err.Error())
   }
}


func principalFunction() {

	var numberProducts int 
	var id int
	var price float64
	var amount int

	fmt.Scan(&numberProducts)

	products := make([]Product, numberProducts)

	for i := 0; i < numberProducts; i++ {

	    fmt.Scan(&id)
	    fmt.Scan(&price)
	    fmt.Scan(&amount)

	    product, err := serializeProduct(id, amount, price)
	    if err != nil {
		fmt.Println(err.Error())
	    }

	    products[i] = product
	}

	msg := detail(products, numberProducts)
	saveProducts(msg)
}
