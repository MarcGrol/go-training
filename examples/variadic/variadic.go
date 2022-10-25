package main

import "fmt"

// START OMIT
func sum(nums ...int) { // HL
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum()    // HL
	sum(1, 2)    // HL
	sum(1, 2, 3) // HL

	nums := []int{1, 2, 3, 4}
	sum(nums...) // HL
}

// END OMIT
