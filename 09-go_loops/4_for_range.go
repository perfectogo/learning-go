package main

import "fmt"

func main() {
	rangeArray()
	rangeString()
}
func rangeArray() {
	var numbers [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numbers)
	for _, number := range numbers {
		fmt.Println(number)
	}
}
func rangeString() {
	var word string = "Go loops"
	fmt.Println(word)
	for _, char := range word {
		fmt.Printf("%c\n", char)
	}
}
