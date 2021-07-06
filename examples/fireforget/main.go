package main

import (
	"fmt"
	"log"
	"time"
)

//START OMIT

func main() {
	a := 1
	b := 1004223
	go func() { // this block runs in background within a go-routine // HL
		result := doWork1(a, b)
		result = doWork2(result)
		result = doWork3(result)
		log.Printf("Result:%s\n", result)
	}()
	log.Printf("Continue without waiting for result\n")

	time.Sleep(time.Second * 5) // Why is this needed?
	log.Printf("main terminates\n")
}

//END OMIT

func doWork1(a, b int) string {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%d", a*b)
}

func doWork2(in string) string {
	time.Sleep(time.Second * 1)
	return "(" + in + ")"
}

func doWork3(in string) string {
	time.Sleep(time.Second * 1)
	return "-->" + in + "<--"
}
