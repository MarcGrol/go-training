package main

import "fmt"

// START OMIT
func sum(a []int, resultChannel chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	resultChannel <- sum // send result back over channel // HL
}

func doit() {
	responseChannel := make(chan int)
	defer close(responseChannel) // prevent resource leak

	go sum([]int{1, 2, 3}, responseChannel)      // 1 + 2 + 3 = 6 // HL
	go sum([]int{4, 5, 6}, responseChannel)      // 4 + 5 + 6 = 15 // HL
	x, y := <-responseChannel, <-responseChannel // receive from channel // HL

	fmt.Printf("one=%d\nanother=%d", x, y) // order undefined
}

// END OMIT

func main() {
	doit()
}
