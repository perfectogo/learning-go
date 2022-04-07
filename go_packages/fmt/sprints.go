package main

import "fmt"

type Student struct {
	name string
	age  int
}

var student Student

func main() {
	student := Student{}
	fmt.Scanf("%s\n", &student.name)
	fmt.Scanf("%d", &student.age)

	std := fmt.Sprintf("name: %s, age: %d", student.name, student.age)
	fmt.Println(std)
}
