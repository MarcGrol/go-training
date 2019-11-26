package main

import "fmt"

// START OMIT
func unhex(c byte) byte {
	switch { // HL
	case '0' <= c && c <= '9': // HL
		return c - '0'
	case 'a' <= c && c <= 'f': // HL
		return c - 'a' + 10
	case 'A' <= c && c <= 'F': // HL
		return c - 'A' + 10
	}
	return 0
}

func main() {
	fmt.Printf("1 -> %d\n", unhex('1'))
	fmt.Printf("9 -> %d\n", unhex('9'))
	fmt.Printf("A -> %d\n", unhex('A'))
	fmt.Printf("F -> %d\n", unhex('F'))
}

// END OMIT
