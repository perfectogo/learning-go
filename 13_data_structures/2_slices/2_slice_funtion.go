package main

import (
	"fmt"
	"reflect"
)

func main() {
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}
	fmt.Println(numbers)

	// copy
	var numbers1 []int = []int{11, 22, 44}
	copy(numbers1, numbers)
	fmt.Println(numbers1)

	// equal
	fmt.Println((reflect.DeepEqual(numbers, numbers1)))

	//len
	fmt.Println(len(numbers))
}
