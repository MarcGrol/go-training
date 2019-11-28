package main

import "log"

// START OMIT
type Person struct { // public
	Name     string  // public
	children []child // private
}

type child struct { // private
	Name string // public
}

func main() {
	me := Person{
		Name: "Marc Grol",
		children: []child{
			{Name: "Pien"},
			{Name: "Tijl"},
		},
	}
	log.Printf("me:%+v", me)
}

// END OMIT
