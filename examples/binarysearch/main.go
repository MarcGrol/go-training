package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 11, -5, 8, 2, 0, 12}
	sort.Ints(numbers)
	fmt.Println("Sorted:", numbers)

	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= 7
	})
	fmt.Printf("The first number >= 7 is at %d and has value %d", index, numbers[index])
}
