package main

import "fmt"

func main() {

	var a, b = 4, 3

	a += b
	fmt.Println(a) //a=7

	a -= b
	fmt.Println(a) //a=4

	a *= b
	fmt.Println(a) //a=12

	a /= b
	fmt.Println(a) //a=4

	a %= b
	fmt.Println(a) //a=1
}
