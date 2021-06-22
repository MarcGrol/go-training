package main

import (
	"fmt"
)

// START OMIT
type Color int // HL

const ( // HL
	Red   Color = iota // 0 (=default) // HL
	Green              // 1 // HL
	Blue               // 2 // HL
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
	var myColor Color // implicily uses default (=Red)
	otherColor := Green
	fmt.Printf("my-color: %v (%d), other-color: %v (%d)\n",
		myColor, myColor, otherColor, otherColor)
}

// END OMIT
