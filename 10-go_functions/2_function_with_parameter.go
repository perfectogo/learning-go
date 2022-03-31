package main

import (
	"fmt"
)

// Return Value from Go Function
func add(num1, num2 int) int {
	return num1 + num2
}

// Return Multiple Values from Go Function
func getInfo(name string, year int) (string, int) {
	return name, 2022 - year
}

func main() {
	result := add(3, 4)
	fmt.Println(result)

	name, age := getInfo("Jasur", 1999)
	fmt.Println(name, age)
}
