package main

import "fmt"

// START OMIT
// This function will be running in the background  // HL
func sum(a []int, resultChannel chan<- int) { // chan<-: Only allows writing to
	sum := 0
	for _, v := range a {
		sum += v
	}
	resultChannel <- sum // send result back over channel
}

func doit() {
	// Create result channel // HL
	responseChannel := make(chan int) // construct channel
	defer close(responseChannel)      // prevent resource leak

	// Divide the work over multiple go-routines that run in background // HL
	go sum([]int{1, 2, 3}, responseChannel) // 1 + 2 + 3 = 6
	go sum([]int{4, 5, 6}, responseChannel) // 4 + 5 + 6 = 15

	// Waitfor all background tasks to have completed 	// HL
	x, y := <-responseChannel, <-responseChannel // receive from channels

	// Continue with result // HL
	fmt.Printf("sum=%d", x+y) // order undefined
}

// END OMIT

func main() {
	doit()
}
