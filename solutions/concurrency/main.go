package main

import (
	"fmt"
	"time"

	"github.com/MarcGrol/go-training/examples/slowapi"
)

func generator(numTasks int) <-chan int {
	responseChannel := make(chan int)

	for i := 0; i < numTasks; i++ {
		go func(idx int) {
			responseChannel <- slowapi.Sum(idx, idx)
		}(i)
	}
	return responseChannel
}

func waitforCompletion(responseChannel <-chan int, taskCount int) (int, int) {
	// only half of the tasks should be completed in 1 sec
	terminationChannel := time.After(10 * time.Second)

	responseCount := 0
	sum := 0
	for {
		select { // blocks until msg received on one of its channels
		case value := <-responseChannel:
			responseCount++
			sum += value
			if responseCount >= taskCount {
				return sum, responseCount
			}
		case <-terminationChannel:
			return sum, responseCount
		}
	}
}

func main() {
	const taskCount = 10000
	responseChannel := generator(taskCount)
	sum, responseCount := waitforCompletion(responseChannel, taskCount)
	fmt.Printf("Got sum %d based on %d responses\n", sum, responseCount)
}
