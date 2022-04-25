package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// START OMIT

func main() {
	fmt.Println(BiggerFloat64(5.1, 1.3))
	fmt.Println(BiggerInt(101, 23))
	fmt.Println(BiggerString("apple", "banana"))

	fmt.Println(BiggerGeneric(5.1, 1.3))
}

// BiggerGeneric takes custom type parameter T of type constraints.Ordered: any type that is order-able
// Ordered is a constraint that permits any ordered type: any type
// that supports the operators < <= >= >.
func BiggerGeneric[T constraints.Ordered](x T, y T) T {
	if x < y {
		return x
	}
	return y
}

// END OMIT

func BiggerFloat64(x float64, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func BiggerInt(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func BiggerString(x string, y string) string {
	if x > y {
		return x
	}
	return y
}
