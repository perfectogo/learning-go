package main

import "fmt"

func main() {
	var number1 int = 7
	var number2 int = 7

	var ptr = new(int)
	*ptr = number1
	fmt.Println(ptr)
	fmt.Println(*ptr)

	ptr = &number2
	fmt.Println(ptr)
	fmt.Println(&ptr)
}
