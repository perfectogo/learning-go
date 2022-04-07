package main

import (
	"fmt"
	"math"
)

func main() {
	//find square root
	fmt.Println(math.Sqrt(36))

	//find cube root
	fmt.Println(math.Cbrt(27))

	fmt.Println(math.Pow(6, 2))

	//
	fmt.Println(math.Max(12, 34))
	fmt.Println(math.Min(12, 34))

	fmt.Println(math.Abs(-234))
	fmt.Println("...")
}
