package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func main() {
	var person1 Person
	fmt.Println(person1)

	//// define the value of name and age
	person1 = Person{"John", 23}
	fmt.Println(person1)
}
