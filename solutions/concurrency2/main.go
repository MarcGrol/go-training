package main

import (
	"fmt"
	"github.com/MarcGrol/go-training/examples/slowapi"
	"sync"
)

func doit(numTasks int) int {
	wg := sync.WaitGroup{}

	allResults := make([]int, numTasks)
	for i := 0; i < numTasks; i++ {
		wg.Add(1)
		go func(idx int) {
			allResults[idx] = slowapi.Sum(idx, idx)
			wg.Done()
		}(i)
	}
	wg.Wait()

	return calculateSum(allResults)
}

func calculateSum(allResults []int) int {
	sum := 0
	for _, v := range allResults {
		sum += v
	}
	return sum
}

func main() {
	const taskCount = 10000

	sum := doit(taskCount)

	fmt.Printf("Got sum %d\n", sum)
}
