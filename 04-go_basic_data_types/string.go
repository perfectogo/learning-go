package main

import "fmt"

func main() {
	var (
		word1 string
		word2 = "world!"
	)

	word1 = "hello"

	// '+' - concatinat
	fmt.Println(word1 + ", " + word2)
}
