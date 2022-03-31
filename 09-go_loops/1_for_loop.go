package main

import "fmt"

func main() {
	var number, fuct = 5, 1

	for i := 1; i <= number; i++ {
		fuct *= i // sum = sum + i
	}

	fmt.Printf("%d! = %d", number, fuct)
}
