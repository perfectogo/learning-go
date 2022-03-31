package main

import "fmt"

func abs(number *int) {
	if *number < 0 {
		*number = -(*number)
	}
}
func main() {
	number := -7
	abs(&number)
	fmt.Println(number)
}
