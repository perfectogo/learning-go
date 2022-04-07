// Running two goroutines concurrently

package main

import (
	"fmt"
	"time"
)

func fn(message string) {

	// infinite for loop
	for {
		fmt.Println(message)

		// to sleep the process for 1 second
		time.Sleep(time.Second * 1)
	}
}

func main() {

	go fn("Process 1")

	fn("Process 2")
}
