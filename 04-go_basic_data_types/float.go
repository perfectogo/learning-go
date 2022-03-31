package main

import "fmt"

func main() {
	var (
		f32 float32 = 6.6 // 32 bits (4 bytes)
		f64 float64       // 64 bits (8 bytes)
	)
	fmt.Println(f32)

	// change value
	f64 = 7.7
	fmt.Println(f64)
}
