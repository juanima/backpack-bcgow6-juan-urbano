package main

import (
    "fmt"
)

const (
    CategoryA = "A"
    CategoryB = "B"
    CategoryC = "C"
    PriceA = 3_000
    PriceB = 1_500
    PriceC = 1_000
    PercentBonoA = 0.5 
    PercentBonoB = 0.2 
)


func computeCategoryC(minutes int) float64 {
    return float64(minutes * PriceC)
}

func computeCategoryB(minutes int) float64 {
    preSalary := float64(minutes) * PriceB
    preSalary = preSalary + (preSalary * PercentBonoB)

    return preSalary
}

func computeCategoryA(minutes int) float64 {
    preSalary := float64(minutes) * PriceA
    preSalary = preSalary + (preSalary * PercentBonoA)

    return preSalary
}

func getComputeCategory(category string) func(int) float64 {
    switch category {
    case CategoryA:
	return computeCategoryA 
    case CategoryB:
	return computeCategoryB 
    case CategoryC:
	return computeCategoryC 
    }
    
    return nil
}


func principalFunction() {

    var minutes int 
    var category string

    fmt.Scan(&category)
    fmt.Scan(&minutes)

    var returnedOperationFunction func(int) float64
    
    returnedOperationFunction = getComputeCategory(category)

    fmt.Println(returnedOperationFunction(minutes))
}

