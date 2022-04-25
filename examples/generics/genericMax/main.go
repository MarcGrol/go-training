package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// START OMIT

func main() {
	fmt.Println(MaxFloat32(5.311, 5.323))
	fmt.Println(MaxFloat64(5.302941, 5.3209943))
	fmt.Println(MaxInt(-101, 23))
	fmt.Println(MaxInt32(10231, 23234))
	fmt.Println(MaxInt64(10234231, 124))
	fmt.Println(MaxUint(101, 23))

	fmt.Println(Max(5.1, 1.3))
}

// Max takes custom type parameter T of type constraints.Ordered: any type that is order-able
// Ordered is a constraint that permits any ordered type:
// any type that supports the operators < <= >= >.
// constraints.Ordered is still in experimental phase
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

func MaxFloat32(x float32, y float32) float32 {
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

func MaxInt32(x int32, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func MaxInt64(x int64, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func MaxUint(x uint, y uint) uint {
	if x > y {
		return x
	}
	return y
}
