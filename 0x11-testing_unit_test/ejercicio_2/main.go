package main

import (
	"fmt"
	"github.com/ejercicio_2/calculator"
)

func main() {
        arr := []int{30, 13, 8, 9}
	sortArr := calculator.SortSlice(arr)
	fmt.Println(sortArr)
}

