package main

import (
	"fmt"
)

func main() {
	var (
		ui8  uint8  // 8 bits (1 byte)
		ui16 uint16 // 16 bits (2 bytes)
		ui32 uint32 // 32 bits (4 bytes)
		ui64 uint64 // 64 bits ( 8 bytes)
		ui   uint   // 32 bits (4 bytes) or 64 bits (8 bytes)
		i    int    // 32 bits (4 bytes) or 64 bits (8 bytes)
		i64  int64  // 64 bits ( 8 bytes)
		i32  int32  // 32 bits (4 bytes)
		i16  int16  // 16 bits (2 bytes)
		i8   int8   // 8 bits (1 byte)
	)
	fmt.Println(ui8, ui16, ui32, ui64, ui)
	fmt.Println(i8, i16, i32, i64, i)

	// change value
	ui, i = 123, 123
	fmt.Printf("%T, %T", ui, i)
	//fmt.Println(ui + i)//ERROR EXPRESSION
}
