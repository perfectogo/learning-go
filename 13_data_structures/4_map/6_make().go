package main

import "fmt"

func main() {
	words := make(map[string]string)
	fmt.Println(words)

	words["go"] = "bormoq"
	words["school"] = "maktab"

	/* words = map[string]string{
		"pen":    "runchka",
		"pencil": "qalam",
		"":       "",
	} */
	fmt.Println(words)
}
