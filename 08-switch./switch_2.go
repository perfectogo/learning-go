package main

import (
	"fmt"
)

func main() {
	switch dayOfWeek := "Sunday"; dayOfWeek {

	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}
}
