package main

import "fmt"

// START OMIT
type Person struct {
	Name        string
	ShirtNumber int
}

// type Stringer interface {
//    String() string
// }
func (p Person) String() string { // HL
	return fmt.Sprintf("%s (number %d)", p.Name, p.ShirtNumber) // HL
} // HL

func main() {
	mj := Person{Name: "Michael Jordan", ShirtNumber: 23}
	jc := Person{Name: "Johan Cruyff", ShirtNumber: 14}
	fmt.Printf("%+v\n%s", mj, jc)
}

// END OMIT
