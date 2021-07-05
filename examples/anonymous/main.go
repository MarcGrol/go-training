package main

import "log"

// START OMIT
func main() {
	persons := []struct { // HL
		name string // HL
		age  int    // HL
	}{ // HL
		{name: "Marc", age: 50},  // HL
		{name: "Eva", age: 48},   // HL
		{name: "Pien", age: 18},  // HL
		{name: "Tijl", age: 15},  // HL
		{name: "Freek", age: 11}, // HL
	} // HL

	for idx, p := range persons {
		log.Printf("person %d: %+v", idx, p)
	}
}

// END OMIT
