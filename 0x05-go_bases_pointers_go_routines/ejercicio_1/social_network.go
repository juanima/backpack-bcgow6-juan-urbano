package main

import (
    "fmt"
    "github.com/juanima/ejercicio_1/models"
)


func detail(users []*models.User) {

    for _, ptrUser := range users {
	user := *ptrUser

	fmt.Printf("%p\n", &user)
	fmt.Printf("%v", user.PrintUser())
	user.SetEmail("patrick@gmail.com")
	fmt.Printf("%v", user.PrintUser())
    }
}


func principalFunction() {

    var ( 
	numberUsers int
        name string 
        lastName string
        email string
	age int
        password string
    )

    fmt.Scan(&numberUsers)
    users := make([]*models.User, numberUsers)

    for i := 0; i < numberUsers; i++ {

	// New line \n to next read line
	fmt.Scanf("%s %s %s %s %d\n", &name, &lastName, &email, &password, &age)

	user, err := models.NewUser(name, lastName, email, password, age) 
	if err != nil {
	    fmt.Println(err.Error())
	}

	users[i] = user
    }

    detail(users)
}
