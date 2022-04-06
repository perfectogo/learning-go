package main

import "fmt"

func main() {

	numbers := [7]int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d -> %d\n", i, numbers[i])
	}
	fmt.Println("\n")

	for index, number := range numbers {
		fmt.Printf("%d -> %d\n", index, number)
	}
}
