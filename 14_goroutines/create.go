package main

import "fmt"

func fn() {
	fmt.Println("process 1")
}
func main() {
	go fn()
	fmt.Println("process 2")
}
