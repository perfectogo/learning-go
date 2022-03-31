package main

import "fmt"

func main() {
	var idf1 bool

	idf1 = (4 < 4) && (7 == 7)
	idf2 := (4 < 4) || (7 == 7)
	idf3 := !("salom" == "salom")

	fmt.Println(idf1)
	fmt.Println(idf2)
	fmt.Println(idf3)
}
