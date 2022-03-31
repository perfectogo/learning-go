package main

import "fmt"

func main() {
	var number1 int = 7
	var number2 int = 7
	//create pointer variablr with *
	var ptr *int = &number1
	fmt.Println(ptr)
	fmt.Println(*ptr)

	ptr = &number2
	fmt.Println(ptr)
}
