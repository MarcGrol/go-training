package main

import (
	"log"
	"sort"
)

// START OMIT
type Person struct {
	name string
	age  int
}

func main() {
	persons := []Person{
		{name: "Marc", age: 50},
		{name: "Pien", age: 18},
		{name: "Eva", age: 48},
		{name: "Tijl", age: 15},
		{name: "Freek", age: 11},
	}

	// END OMIT
	// sort in age "ascending"
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].age < persons[j].age
	})
	log.Printf("%+v", persons)
}
