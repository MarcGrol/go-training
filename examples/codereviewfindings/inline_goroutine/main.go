package main

import (
	"fmt"
	"log"
	"time"
)

// START OMIT
func inline_goroutine(input []int) {
	for _, val := range input {
		go func() {
			doSomething(val)
		}()
	}
}

// END OMIT

func main() {
	input := []int{1, 2, 3, 4, 5}

	inline_goroutine(input)
	time.Sleep(1 * time.Second)
}

func doSomething(in int) string {
	val := fmt.Sprintf("%d", in)
	log.Printf("%s", val)
	return val
}
