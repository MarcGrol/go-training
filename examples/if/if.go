package main

import (
	"fmt"
)

func main() {
	num := 9
	if num < 0 { // HL
		fmt.Println(num, "is negative")
	} else if num < 10 { // HL
		fmt.Println(num, "has 1 digit")
	} else { // HL
		fmt.Println(num, "has multiple digits")
	}
}
