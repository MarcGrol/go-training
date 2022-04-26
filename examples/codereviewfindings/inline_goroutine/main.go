package main

import (
	"fmt"
	"log"
	"time"
)

// START OMIT
func wrong(input []int) {
	for _, val := range input {
		go func() {
			doSomething(val)
		}()
	}
}

func correct(input []int) {
	for _, val := range input {
		go func(in int) {
			doSomething(in)
		}(val)
	}
}

// END OMIT

func main() {
	input := []int{1, 2, 3, 4, 5}

	log.Printf("Wrong\n")
	wrong(input)
	time.Sleep(1 * time.Second)

	log.Printf("Success:\n")
	correct(input)
	time.Sleep(1 * time.Second)

}

func doSomething(in int) string {
	val := fmt.Sprintf("%d", in)
	log.Printf("%s", val)
	return val
}
