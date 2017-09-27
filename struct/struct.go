package main

import "fmt"

var me Person = Person{
	Name:      "Marc Grol",
	Interests: []string{"Running", "Golang"},
	Children: []Child{
		{Name: "Pien", Age: 12},
		{Name: "Tijl", Age: 9},
	},
}

// START OMIT
type Person struct {
	Name      string
	Interests []string
	Children  []Child
}

type Child struct {
	Name string
	Age  int
}

func (p *Person) AddChild(child Child) { // HL
	p.Children = append(p.Children, child)
}

func main() {
	// me := Person{...}
	me.AddChild(Child{Name: "Freek", Age: 5})
	fmt.Printf("%+v\n", me) // Note: debug struct with %+v
}

// END OMIT
