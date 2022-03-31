package main

import "fmt"

func greet(name *string) *string {
	if *name != "" {
		*name = "hello! " + *name
		return name
	} else {
		msg := "no name"
		return &msg
	}
}
func main() {
	name := ""
	greet := greet(&name)
	fmt.Println(*greet)
}
