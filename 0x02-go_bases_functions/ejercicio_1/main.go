package main

import "fmt"

func main() {
	var queries int
	var line string

	fmt.Scan(&queries)

	for queries > 0 {
		fmt.Scan(&line)
		fmt.Println(line)

		principalFunction()

		queries -= 1
	}
}


