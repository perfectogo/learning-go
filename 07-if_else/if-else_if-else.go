package main

import "fmt"

func main() {
	number := 777

	if (0 <= number) && (number < 10) {
		fmt.Println(1)
	} else if (9 < number) && (number < 100) {
		fmt.Println(2)
	} else if (99 < number) && (number < 1000) {
		fmt.Println(3)
	} else {
		fmt.Println(404)
	}
}
