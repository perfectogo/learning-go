package main

import "fmt"

func main() {
	var number int = 5

	// anonymous function 1
	fmt.Println(func() int {
		var fuct int = 1
		for i := 1; i <= number; i++ {
			fuct *= i
		}
		return fuct
	}())

	// anonymous function 2
	add := func(num1, num2 int) {
		fmt.Println(num1 + num2)
	}
	//clouser
	add(3, 4)
}
