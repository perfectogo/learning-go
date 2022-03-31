package main

import "fmt"

const (
	Sunday uint8 = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Wednesday)
	fmt.Println(Saturday)
}
