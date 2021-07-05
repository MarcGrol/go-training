package main

import (
	"fmt"
)

// START OMIT
type Parent struct {
	Name      string
	Interests []string
	Children  []Child
}

type Child struct {
	Name string
	Age  int
}

func main() {
	parent := &Parent{
		Name: "Marc",
		Children: []Child{
			{Name: "Pien", Age: 18},
		},
	}
	fmt.Printf("%+v", parent)
}

// END OMIT
