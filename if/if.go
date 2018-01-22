package main

import (
	"fmt"
)

// START OMIT
func main() {
	num := 9
	if num < 0 { // HL
		fmt.Println(num, "is negative")
	} else if num < 10 { // HL
		fmt.Println(num, "has 1 digit")
	} else { // HL
		fmt.Println(num, "has multiple digits")
	}

	if smallerNum := minus10(num); smallerNum > 0 {  // HL
		fmt.Println(smallerNum, " is larger than 0")
	} else {
		fmt.Println(smallerNum, " is smaller than 0")
	}
}

func minus10(a int) int {
	return a - 10
}
// END OMIT
