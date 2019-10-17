package main

import "fmt"

func main() {
	// START OMIT
	sum := 1
	for sum < 1000 { // HL
		sum += sum
	}
	fmt.Println(sum)
	// END OMIT
}
