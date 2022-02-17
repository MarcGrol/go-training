package main

import (
	"fmt"

	"github.com/MarcGrol/go-training/solutions/flow/calclib"
)

func main() {
	// Sum of all values from 1 up to 100
	{
		sum := 0
		for i := 0; i <= 100; i++ {
			sum += i
		}
		fmt.Printf("sum[0..100] = %d\n", sum)
	}

	{
		sum := 0
		iterCount := 0
		for i := 0; ; i++ {
			sum += i
			if sum >= 1000 {
				break
			}
			iterCount++
		}
		fmt.Printf("nu-iteration (>=1000) = %d\n", iterCount)
	}

	iterCount := calclib.SumUntilMax(1000)
	fmt.Printf("num iterations [0..100] = %d\n", iterCount)
}
