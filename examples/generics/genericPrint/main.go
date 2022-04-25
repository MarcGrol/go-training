package main

import "fmt"

// START OMIT

type CustomInt int
type CustomString string

type CustomTypes interface {
	string | float64 | ~int
}

func main() {
	print("string")
	print(1.1)
	print(1)
	print(CustomInt(123))

	//print(CustomString("Custom string"))
}

func print[T CustomTypes](value T) {
	fmt.Println(value)
}

// END OMIT
