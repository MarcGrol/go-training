package main

import "fmt"

func main() {
	// START OMIT
	values := []string{"a", "b", "c"}
	for idx, value := range values { // HL
		fmt.Printf("%d:%s\n", idx, value)
	}
	// END OMIT
}
