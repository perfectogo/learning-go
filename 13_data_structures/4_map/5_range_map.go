package main

import "fmt"

func main() {

	// create with shorthand
	langSt := map[string]float32{
		"Golang": 85,
		"Java":   80,
		"Python": 81,
	}

	for lang, status := range langSt {

		fmt.Printf("language: %s, status: %f\n", lang, status)
	}

}
