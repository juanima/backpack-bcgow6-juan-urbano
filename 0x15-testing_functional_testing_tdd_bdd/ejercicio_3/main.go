package main

import (
	"fmt"

	"github.com/ejercicio_3/secuence"
)

func main() {

        number := 5
	sucesion, total := secuence.Fibonnaci(number)

        fmt.Println("n-ésimo número", number)
	fmt.Println("Sucesión ", sucesion)
	fmt.Println("Total ", total)
}
