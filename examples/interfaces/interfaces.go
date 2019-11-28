package main

import "fmt"

// START OMIT
type Sporter struct {
	Name        string
	ShirtNumber int
}

// type Stringer interface {
//    String() string
// }
func (p Sporter) String() string { // HL
	return fmt.Sprintf("%s (number %d)", p.Name, p.ShirtNumber) // HL
} // HL

func main() {
	mj := Sporter{Name: "Michael Jordan", ShirtNumber: 23}
	jc := Sporter{Name: "Johan Cruyff", ShirtNumber: 14}
	fmt.Printf("%+v\n%s", mj, jc)
}

// END OMIT
