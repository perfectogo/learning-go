package main

import "fmt"

func main() {
	var number int

	//get input values from the user
	fmt.Scan(&number)
	fmt.Println(number)

	//get input values using the format specifier
	fmt.Scanf("%d", &number)
	fmt.Println(number)
}
