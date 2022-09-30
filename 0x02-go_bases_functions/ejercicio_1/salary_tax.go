package main

import "fmt"

func computeTaxes(salary float64) float64 {

	if salary > 150_000 {
		return salary * 0.1
	} else if salary > 50_000 && salary <= 150_000 {
		return salary * 0.17
	}

	return 5
}

func principalFunction() {

	var salary float64

	fmt.Scan(&salary)

	fmt.Printf("El impuesto del salario es $%.3f USD\n", computeTaxes(salary))
}
