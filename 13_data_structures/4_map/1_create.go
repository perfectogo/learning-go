package main

import "fmt"

func main() {
	// declare for type
	var mapp map[string]string
	// we can't use map befor make

	//mapp["salom"] = "hello" // ERROR
	fmt.Println(mapp)

	// create with shorthand
	langSt := map[string]float32{
		"Golang": 85,
		"Java":   80,
		"Python": 81,
	}

	fmt.Println(langSt)

	// Access Values of a Map
	fmt.Println(langSt["Golang"])

	// Change value of a map
	langSt["Golang"] = 100
	fmt.Println(langSt["Golang"])

}
