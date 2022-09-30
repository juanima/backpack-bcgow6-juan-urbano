package main

import (
    "fmt"
    "errors"
    "encoding/json"
)

type Student struct {
    Name string `json:"name"`
    LastName string `json:"lastname"`
    Dni int `json:"dni"`
    Date string `json:"date"`
}

func (s Student) Detail() {
    fmt.Println("Guau!")
}


func serializeStudent(name string, lastName string, date string, dni int) (studentJSON []byte, err error) {

    student := Student {
	Name: name, 
	LastName: lastName, 
	Dni: dni, 
	Date: date, 
    }

    student.Detail()
    studentJSON, err = json.Marshal(&student)
    if err != nil {
	fmt.Print(err.Error())
	return nil, errors.New("Panico")
    }

    return studentJSON, nil
}


func principalFunction() {

	var numberStudents int 
	var name string
	var lastName string
	var dni int
	var date string

	fmt.Scan(&numberStudents)

	for i := 0; i < numberStudents; i++ {

	    fmt.Scan(&name)
	    fmt.Scan(&lastName)
	    fmt.Scan(&dni)
	    fmt.Scan(&date)

	    student, err := serializeStudent(name, lastName, date, dni)
	    if err != nil {
		fmt.Println(err.Error())
	    }

	    fmt.Printf("Info del estudiante %s \n", string(student))
	}
}
