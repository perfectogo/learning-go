package main

import "fmt"

func main() {
	// declare
	var word string = "hello"
	fmt.Println(word)
	// concatenates
	sentences := word + ", " + "world!"

	// print with len
	fmt.Println(sentences, len(sentences))
}
