package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) modify(name string) {
	p.Name = name
}

func (p Person) thisModifyDoesNotWork(name string) {
	p.Name = name
}

func passByValue(person Person) {
	person.Name = "Silvan"

}

func passByReference(person *Person) {
	person.Name = "Silvan"
}

func main() {
	person := Person{Name: "Marc"}

	passByValue(person)
	fmt.Printf("by value: not adjusted: %+v\n", person)
	// results in by value: not adjusted: {Name:Marc}

	passByReference(&person)
	fmt.Printf("by reference adjusted: %+v\n", person)
	// results in: by reference adjusted: {Name:Silvan}

	person.thisModifyDoesNotWork("Henk")
	fmt.Printf("by value: not adjusted: %+v\n", person)
	// results in: by value: not adjusted: {Name:Silvan}

	person.modify("Henk")
	fmt.Printf("by reference adjusted: %+v\n", person)
	// results in: by reference adjusted: {Name:Henk}

}
