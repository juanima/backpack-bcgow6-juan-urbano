package main


import (
    "fmt"
    "github.com/ejercicio_1/internal/service"
    "github.com/ejercicio_1/internal/file"
)

const (
    PATH = "go-web/products.json"
)


func main() {

        var (
	    products []service.Product
	    data    file.File
	)

	data.SetPath(PATH)
	products, err := data.Read()
	if err != nil {
	    panic(err.Error())
	}
														    
	fmt.Printf("%+v\n", products[0])
	fmt.Printf("%+v\n", products[1])
}
