// hello.go
package main

import (
	"fmt"
)

type Person struct {
	Name      string
	Interests []string
	Children  []Child
}

type Child struct {
	Name string
	Age  int
}

// START OMIT
func NewPerson(name string, interests ...string) *Person {
	person := new(Person) // HL
	person.Name = name
	person.Interests = make([]string, 0, len(interests)) // HL
	for _, interest := range interests {
		person.Interests = append(person.Interests, interest)
	}
	person.Children = make([]Child, 0, 10) // HL
	return person
}

func main() {
	me := Person{
		Name:      "Marc Grol",
		Interests: []string{"Running", "Golang"},
		Children: []Child{
			{Name: "Pien", Age: 12},
			{Name: "Tijl", Age: 9},
			{Name: "Freek", Age: 5},
		},
	}
	you := NewPerson("Eva Berkhout", "Running", "Reading")
	fmt.Printf("me:%+v\nyou:%+v", me, you)
}

// END OMIT
