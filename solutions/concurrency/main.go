package main

import (
	"fmt"
	"time"
)

func SlowActionWithChannel(a,b int, responseChannel chan int) {
	responseChannel <- SimulateSlowAction(a,b)
}

func waitforCompletion(responseChannel chan int) int {
	terminationChannel := time.After(1 * time.Second)

	responseCount := 0
	for {
		select { // blocks until msg received on one of its channels
		case <-responseChannel:
			responseCount++
			if responseCount >= 100 {
				return responseCount
			}
		case <-terminationChannel:
			return responseCount // break out of loop
		}
	}
}

func main() {
	responseChannel := make(chan int)
	defer close(responseChannel)

	for i:=0; i<100; i++ {
		go SlowActionWithChannel(i,i,responseChannel)
	}

	responseCount := waitforCompletion(responseChannel)
	fmt.Printf("Got %d responses\n", responseCount)
}