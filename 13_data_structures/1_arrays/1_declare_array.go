package main

import "fmt"

func main() {
	// declare
	var numbers1 [10]int
	numbers1 = [10]int{1, 2, 4, 5, 6}
	fmt.Println(numbers1)
	fmt.Println(numbers1[0])

	// declare with [size]type
	var numbers2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers2)
	fmt.Println(numbers2[3])

	// declare without [size]type
	var numbers3 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers3)

	// declare shorthand
	numbers4 := numbers1
	fmt.Println(numbers4)
}
