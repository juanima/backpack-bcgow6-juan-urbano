package main

import (
    "fmt"
    "errors"
)

func averageGrade(grades ...float64) (float64, error) {

    var result float64
    var numberGrades int

    for i, value := range grades {

	if value < 0 {
	    return 0, errors.New("La nota debe ser debe ser un número positivo")    
	    //continue
	}
	    
	result += value
	numberGrades = i
    }

    numberGrades++

    return result / float64(numberGrades), nil
}

func principalFunction() {

	var numberGrade int 

	fmt.Scan(&numberGrade)

	grades := make([]float64, numberGrade)

	for i := 0; i < numberGrade; i++ {
	    fmt.Scan(&grades[i])
	}

	average, err := averageGrade(grades...)
	if err != nil {
	    fmt.Println("Hubo un error")
	}

	fmt.Printf("El promedio calculado para el estudiante es %.2f\n", average)
}
