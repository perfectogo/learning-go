package main

import "fmt"

func main() {
	// explicitly declaring the data type and assigning value
	var v1 int = 1
	// Assigning a value without explicitly declaring the data type
	var v2 = 2

	var (
		v3, v4 int  = 3, 4
		v5     int8 = 4
	)

	// shorthand notation to define variable
	v6 := 5

	fmt.Printf("v1 = %d, v2 = %d, v3 = %d, v4 = %d, v5 = %d, v6 = %d", v1, v2, v3, v4, v5, v6)
}
