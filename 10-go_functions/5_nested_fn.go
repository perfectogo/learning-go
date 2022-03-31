package main

import "fmt"

func fuct(number int) {
	fn := func() {
		fc := 1
		for i := 1; i <= number; i++ {
			fc *= i
		}
		fmt.Println(fc)
	}

	fn()
}
func main() {
	fuct(5)
}
