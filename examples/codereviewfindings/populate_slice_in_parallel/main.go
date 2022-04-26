package main

import (
	"fmt"
	"log"
	"time"
)

// START OMIT
func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := []string{}

	for _, val := range input {
		go func(val int) {
			output = append(output, fmt.Sprintf("%d", val))
		}(val)
	}

	time.Sleep(1 * time.Second)

	// print contents of slice
	for _, s := range output {
		log.Printf("%s\n", s)
	}
}

// END OMIT
