package main

import (
	"fmt"
	"time"

	"github.com/MarcGrol/go-training/examples/slowapi"
)

const (
	taskCount = 100
)

func SlowSummer(a, b int, responseChannel chan int) {
	responseChannel <- slowapi.Sum(a, b)
}

func waitforCompletion(responseChannel chan int) (int, int) {
	// only half of the tasks should be completed in 1 sec
	terminationChannel := time.After(1 * time.Second)

	responseCount := 0
	sum := 0
	for {
		select { // blocks until msg received on one of its channels
		case value := <-responseChannel:
			responseCount++
			sum += value
			if responseCount >= taskCount {
				break
			}
		case <-terminationChannel:
			break
		}
	}
	return sum, responseCount
}

func main() {
	responseChannel := make(chan int)
	defer close(responseChannel)

	for i := 0; i < taskCount; i++ {
		go SlowSummer(i, i, responseChannel)
	}

	sum, responseCount := waitforCompletion(responseChannel)
	fmt.Printf("Got sum %d based on %d responses\n", sum, responseCount)
}
