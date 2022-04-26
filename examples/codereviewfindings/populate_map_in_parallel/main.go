package main

import (
	"fmt"
	"log"
	"time"
)

// START OMIT
func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	output := map[int]string{}

	for _, val := range input {
		go func(val int) {
			output[val] = fmt.Sprintf("%d", val)
		}(val)
	}

	time.Sleep(1 * time.Second)

	// print contents of slice
	for i, s := range output {
		log.Printf("%d: %s\n", i, s)
	}
}

// END OMIT

// reserve slot first
//for _, val := range input {
//  output[val] = ""
//}
