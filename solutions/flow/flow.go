package main

import (
	"fmt"
	"github.com/MarcGrol/go-training/solutions/flow/calclib"
)

func main() {
	// Sum of all values from 1 up to 100
	sum := 0
	for i:=0; i<=100; i++ {
		sum+=i
	}
	fmt.Printf("sum[0..100] = %d\n", sum)

	iterCount := calclib.SumUntillMax(1000)
	fmt.Printf("num iterations [0..100] = %d\n", iterCount)
}


