package main

import (
    "fmt"
    "github.com/juanima/ejercicio_2/models"
)


func detail(users []*models.User) {

    for _, ptrUser := range users {
	user := *ptrUser

	fmt.Printf("%p\n", &user)
	fmt.Printf("%v", user.StrUser())
    }
}

func deleteUser(nameUser, lastName string, users *[]*models.User) {

    user := findUser(nameUser, lastName, users) 
    products := (*user).GetProducts()
    
    // To remove all elements, simply set the slice to nil.
    products = nil
    (*user).SetProducts(&products)
}

func updateProducts(nameUser, lastName, nameProduct string, price float64, amount int, users *[]*models.User) {

    user := findUser(nameUser, lastName, users) 

    product, err := models.NewProduct(nameProduct, price) 
    if err != nil {
	fmt.Println(err.Error())
    }

    product.SetAmount(amount)
    addProduct(&user, product, amount) 
}



func findUser(name string, lastName string, users *([]*models.User)) (user *models.User) {

    for _, user := range *users {
	// fmt.Println("%v\n", user)
	if user.GetName() == name && user.GetLastName() == lastName {
	    return user
	}
    }

    return user
}


func addProduct(user **models.User, newProduct *models.Product, amount int) {

    isUpdate := false 
    products := (*user).GetProducts()

    for _, product := range products {
	if product.GetName() == newProduct.GetName() {
	    product.SetAmount(product.GetAmount() + 1)
	    isUpdate = true
	}
    }
    
    if isUpdate == false {
	copyProducts := make([]*models.Product, len(products) + 1, cap(products) * 2) 
	
	copy(copyProducts, products)
	copyProducts[len(products)] = newProduct

	// fmt.Println("-----")
	// for _, product := range copyProducts {
	//     fmt.Printf(" Product: %s\n", product.GetName())
	// }
	// fmt.Println("-----")

	// fmt.Printf("Before products %p -- %p\n", products, copyProducts)
	(*user).SetProducts(&copyProducts)
    }
}


func principalFunction() {

    var ( 
	numberUsers int
	numberProducts int
        nameUser string 
        lastName string
        email string
	nameProduct string
	price float64
	amount int
    )

    fmt.Scan(&numberUsers)
    users := make([]*models.User, numberUsers)

    for i := 0; i < numberUsers; i++ {

	// New line \n to next read line
	fmt.Scanf("%s %s %s\n", &nameUser, &lastName, &email)

	user, err := models.NewUser(nameUser, lastName, email) 
	if err != nil {
	    fmt.Println(err.Error())
	}
	
	fmt.Scanf("%d", &numberProducts)
	products := make([]*models.Product, numberProducts)

	for j := 0; j < numberProducts; j++ {

	    fmt.Scanf("%s %f %d\n", &nameProduct, &price, &amount)

	    product, err := models.NewProduct(nameProduct, price) 
	    if err != nil {
		fmt.Println(err.Error())
	    }

	    products[j] = product
	}

	user.SetProducts(&products)
	users[i] = user
    }

    detail(users)

    // add a product
    fmt.Scanf("%s %s %s %f %d\n", &nameUser, &lastName, &nameProduct, &price, &amount)
    updateProducts(nameUser, lastName, nameProduct, price, amount, &users) 

    fmt.Println("------")
    detail(users)

    // delete a product 
    fmt.Scanf("%s %s\n", &nameUser, &lastName)
    deleteUser(nameUser, lastName, &users)

    fmt.Println("------")
    detail(users)
}
