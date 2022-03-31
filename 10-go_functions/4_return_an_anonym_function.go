package main

import "fmt"

func fuct(num1, num2 int) func() int {

	sum := num1 + num2

	return func() int {
		var fct = 1
		for i := 1; i <= sum; i++ {
			fct *= i
		}
		return fct
	}
}

func main() {
	fuct := fuct(2, 3)
	fmt.Println(fuct())
}
