package main

import "fmt"

func main() {
	number := 1

	for {

		fmt.Printf("%d\n", number)

		if number > 7 {
			break
		}
		number++
	}
}
