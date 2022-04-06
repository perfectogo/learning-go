package main

import "fmt"

func main() {
	words := make(map[string]string)

	words["go"] = "bormoq"
	words["school"] = "maktab"
	value, ok := words["gol"]
	if ok {
		fmt.Println(value)
	} else {

		fmt.Println("no value")

	}

}
