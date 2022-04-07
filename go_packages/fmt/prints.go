package main

import "fmt"

func main() {
	// print without newline
	fmt.Print("hello")
	fmt.Print("hello\n")

	// print with newline
	fmt.Println("hello, world")

	// print with format
	var word string = "godbye!"
	fmt.Printf("%s, %T\n", word, word)
}
