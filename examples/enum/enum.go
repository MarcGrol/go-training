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
	case Red:
		return "red"
	case Green:
		return "green"
	case Blue:
		return "blue"
	default:
		return "unknown"
	}
}
func main() {
	myColor := Blue
	fmt.Printf("my-color: %d - %s\n", myColor, myColor)
}

// END OMIT
