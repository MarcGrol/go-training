package main

import "fmt"

// START OMIT

type CustomInt int
type CustomString string

type CustomTypes interface {
	string | float64 | ~int
}

func main() {
	printCustomType("string")
	printCustomType(1.1)
	printCustomType(1)
	printCustomType(CustomInt(123))

	//printCustomType(CustomString("Custom string"))
}

// END OMIT

func printCustomType[T CustomTypes](value T) {
	fmt.Println(value)
}
