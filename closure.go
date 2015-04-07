package main

import "fmt"

// START OMIT
func intSeq() func() int {
	i := 0
	return func() int { // HL
		i += 1   // HL
		return i // HL
	} // HL
}

func main() {
	nextInt := intSeq()

	// See the effect of the closure by calling nextInt a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that particular function
	newInts := intSeq()
	fmt.Println(newInts())
}

// END OMIT
