package main

import "fmt"

// START OMIT
func main() {
	sum := 0
	for i := 0; i < 10; i++ { // HL
		sum += i
	}
	fmt.Println(sum)
}

// END OMIT
