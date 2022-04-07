package main

import "fmt"

type Area struct {
	a uint
	b uint
}

func main() {
	area1 := Area{3, 5}
	fmt.Println(area1.a)
	fmt.Println(area1.b)
}
