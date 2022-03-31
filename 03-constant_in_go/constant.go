package main

import "fmt"

func main() {
	const c1 int = 1
	const c2 = 2

	const (
		c3, c4 int  = 3, 4
		c5     int8 = 5
	)
	fmt.Printf("v1 = %d, v2 = %d, v3 = %d, v4 = %d, v5 = %d", c1, c2, c3, c4, c5)
}
