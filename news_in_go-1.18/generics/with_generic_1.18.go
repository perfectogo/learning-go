package main

import(
	"log"
)
type Numeric interface {
    int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Number[typ ~int|~string](elm typ) typ {
	return elm
}


func Join[typ any] (elements ...typ) []typ{
	return elements

}

func main() {
	slice1 :=Join[int](1, 2, 4)
	slice2 :=Join[bool](true, true, true)
	slice3 :=Join[string]("true", "true", "true")
	slice4 :=Number[string]("Hello, World!")

	log.Println(slice1)
	log.Println(slice2)
	log.Println(slice3)
	log.Println(slice4)
}
