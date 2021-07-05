package main

import "fmt"

// START OMIT
func main() {
	var anything interface{} = 1
	fmt.Printf("anything int: %+v\n", anything)

	anything = "marc"
	print(anything)
}

func print(it interface{}) {
	fmt.Printf("anything: %+v\n", it)
}

// END OMIT
