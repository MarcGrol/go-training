package main

/*
#include personlib/person.h
*/
import "C"
import "fmt"

type CPerson struct {
	Person C.Person
}

func main() {
	person := CPerson{
		Person: C.Person{
			id:   "",
			name: "",
			age:  0,
		},
	}

	fmt.Printf("%+v", person)
}
