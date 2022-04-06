package main

import "fmt"

func main() {

	// create with shorthand
	langSt := map[string]float32{
		"Golang": 85,
		"Java":   80,
		"Python": 81,
	}

	fmt.Println(langSt)

	// add Element of Go Map Element
	langSt["C++"] = 90
	fmt.Println(langSt)
}
