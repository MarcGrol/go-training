package main

import (
	"fmt"
)

// START OMIT
type Color int // HL

const ( // HL
	Unknown Color = iota // 0 (=default) // HL
	Red                  // 1 // HL
	Green                // 2 // HL
	Blue                 // 3 // HL
) // HL

func (c Color) String() string {
	switch c {
	case Green:
		return "green"
	case Blue:
		return "blue"
	default:
		return "red"
	}
}
func main() {
	var myColor Color // uses default
	otherColor := Green
	fmt.Printf("my-color: %v (%d), other-color: %v (%d)\n",
		myColor, myColor, otherColor, otherColor)
}

// END OMIT
