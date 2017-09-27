package main

import "fmt"

type Person struct { // implements Stringer interface
	Name        string
	ShirtNumber int
}

// type Stringer interface {
//    String() string
// }
func (p Person) String() string {
	return fmt.Sprintf("%s (number %d)", p.Name, p.ShirtNumber)
}

func main() {
	a := Person{Name: "Michael Jordan", ShirtNumber: 23}
	z := Person{Name: "Johan Cruyff", ShirtNumber: 14}
	fmt.Printf("%+v - %s", a, z)
}
