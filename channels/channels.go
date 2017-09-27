package main

import "fmt"

// START OMIT
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)     // HL
	go sum(a[:len(a)/2], c) // []int{7,2,8} -> 17 // HL
	go sum(a[len(a)/2:], c) // []int{-9,4,0} -> -5 // HL
	x, y := <-c, <-c        // receive from c // HL

	fmt.Printf("one=%d\nanother=%d", x, y) // order undefined
}

// END OMIT
