package main

import "fmt"

func main() {
	fromFloatToInt()
	fmt.Println("\n")
	fromIntToFloat()
}
func fromFloatToInt() {
	var floatValue float32 = 5.45

	// type conversion from float to int
	var intValue int = int(floatValue)

	fmt.Printf("Float Value is %g\n", floatValue)
	fmt.Printf("Integer Value is %d", intValue)
}

func fromIntToFloat() {
	var intValue int = 2

	// type conversion from int to float
	var floatValue float32 = float32(intValue)

	fmt.Printf("Integer Value is %d\n", intValue)
	fmt.Printf("Float Value is %f", floatValue)
}
