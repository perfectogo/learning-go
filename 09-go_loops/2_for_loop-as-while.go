package main

import "fmt"

func main() {
	var (
		number int = 124
		sum        = 0
	)

	for number >= 1 {
		sum += number % 10
		number /= 10
	}

	fmt.Println(sum)
}
