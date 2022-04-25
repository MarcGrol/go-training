package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// START OMIT

func main() {
	fmt.Println(MaxFloat64(5.1, 1.3))
	fmt.Println(MaxInt(101, 23))
	fmt.Println(MaxString("apple", "banana"))

	fmt.Println(Max(5.1, 1.3))
}

// Max takes custom type parameter T of type constraints.Ordered: any type that is order-able
// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
func Max[T constraints.Ordered](x T, y T) T {
	if x > y {
		return x
	}
	return y
}

// END OMIT

func MaxFloat64(x float64, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func MaxInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxString(x string, y string) string {
	if x > y {
		return x
	}
	return y
}
