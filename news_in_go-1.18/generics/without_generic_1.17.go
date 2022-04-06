package main

import (
	"log"
)

func JoinBool(elements ...bool) []bool {
	return elements
}

func JoinInt(elements ...int) []int {
	return elements
}

func JoinString(elements ...string) []string {
	return elements
}

func main() {
	slice1 := JoinInt(1, 2, 3)
	slice2 := JoinBool(true, true, false)
	slice3 := JoinString("a", "b", "c")

	log.Println(slice1)
	log.Println(slice2)
	log.Println(slice3)
}
