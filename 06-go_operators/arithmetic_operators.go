package main

import "fmt"

func main() {
	var a, b = 4, 3

	add := a + b
	sub := b - a
	mul := a * b
	div := a / b
	mod := a % b

	fmt.Println(add)
	fmt.Println(sub)
	fmt.Println(mul)
	fmt.Println(div)
	fmt.Println(mod)
}
