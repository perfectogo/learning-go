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

	// delete Element of Go Map Element
	delete(langSt, "Python")
	fmt.Println(langSt)
}
