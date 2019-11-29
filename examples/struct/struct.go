package main

import "fmt"

// START OMIT
type Student struct { // public
	Name string       // public
	password  string  // private
	teacher teacher   // private
}

type teacher struct { // not accessible outside package
	Name string
}

func main() {
	son := Student{ // constructor like
		Name: "Freek",
		password:  "secret",
		teacher:teacher{
			Name: "Lisette",
		},
	}
	fmt.Printf("%+v", son) // %+v: convenience debugging
}

// END OMIT
