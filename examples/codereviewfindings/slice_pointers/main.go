package main

import "log"

// START OMIT
func main() {
	input := []string{"a", "b", "c"}
	output := []*string{}

	// convert to slice of pointers
	for _, s := range input {
		output = append(output, &s)
	}

	// print contents of slice
	for _, s := range output {
		log.Printf("%p: %s\n", s, *s)
	}
}

// END OMIT
