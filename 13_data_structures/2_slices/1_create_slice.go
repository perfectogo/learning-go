package main

import "fmt"

func main() {
	//creating slice
	var numbers1 = []int{1}
	fmt.Println(numbers1)

	//creating slice with shorthand
	numbers2 := []int{1, 2}
	fmt.Println(numbers2)

	//Create slice from an array
	var numbers3 []int
	var arrayNumbers = [5]int{1, 2, 3, 4, 5}

	numbers3 = arrayNumbers[2:4]
	fmt.Println(numbers3)
}
