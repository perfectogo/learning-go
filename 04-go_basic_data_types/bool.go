package main

import "fmt"

func main() {
	var idf1 bool = true
	var idf2 = false
	idf3 := true
	idf4 := 4 < 5

	fmt.Println(idf1)
	fmt.Println(idf2)
	fmt.Println(idf3)
	fmt.Println(idf4)
	fmt.Println(idf3 || idf2 && idf3)
}
