package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-2) + fibonacci(n-1)
	}
}

func main() {
	fibonacciNumber := fibonacci(7)
	fmt.Println(fibonacciNumber)
}
