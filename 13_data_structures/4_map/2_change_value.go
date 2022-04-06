package main

import "fmt"

func main() {

	langSt := map[string]float32{
		"Golang": 85,
		"Java":   80,
		"Python": 81,
	}

	fmt.Println(langSt["Golang"])

	// Change value of a map
	langSt["Golang"] = 100
	fmt.Println(langSt["Golang"])

}
