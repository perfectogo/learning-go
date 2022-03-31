package main

import (
	"fmt"
)

func main() {
	age := 28

	// switch without any expression
	switch {

	case (1 < age) && (age <= 5):
		fmt.Println("baby")

	case (5 < age) && (age <= 14):
		fmt.Println("child")

	case (14 < age) && (age <= 18):
		fmt.Println("teen")

	case (18 < age) && (age <= 30):
		fmt.Println("young")

	case (30 < age) && (age <= 50):
		fmt.Println("middle age")

	case (50 < age):
		fmt.Println("old")

	default:
		fmt.Println("Not February")
	}
}
