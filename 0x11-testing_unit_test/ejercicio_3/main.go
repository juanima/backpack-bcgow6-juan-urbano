package main

import (
	"fmt"
	"github.com/ejercicio_3/calculator"
)

func main() {
        quotient, divisor:= 10, 5

	dividiend, err := calculator.Divide(quotient, divisor)
        if err != nil {
                fmt.Println("Error present")
        }

	fmt.Println(dividiend)
}
