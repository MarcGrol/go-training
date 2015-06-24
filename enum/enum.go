package main

import (
	"fmt"
)

// START OMIT
type Color int

const (
	Red   Color = iota // 0 (=default)
	Green              // 1
	Blue               // 2
)

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
	myColor := Green
	fmt.Printf("my-color: %d - %s\n", myColor, myColor)
}

// END OMIT
