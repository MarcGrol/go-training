package main

import "fmt"

func main() {
	var anything interface{} = 1
	fmt.Printf("anything int: %+v\n", anything)

	anything = "marc"
	fmt.Printf("anything string: %+v\n", anything)

	anything = false
	print(anything)
}

func print(it interface{}) {
	fmt.Printf("anything: %+v\n", it)
}
