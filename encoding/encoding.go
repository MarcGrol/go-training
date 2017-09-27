// hello.go
package main

import (
	"encoding/json"
	"encoding/xml"
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

func main() {
	var jsonMe string = `
	{
		"Name":"Marc Grol",
	    "Interests":["Running","Golang"],
	    "Children":[
	    	{"Name":"Pien","ShirtNumber":12},
	    	{"Name":"Tijl","ShirtNumber":9},
	    	{"Name":"Freek","ShirtNumber":5}
	    ]
	}`

	var me Person
	json.Unmarshal([]byte(jsonMe), &me) // HL

	fmt.Printf("About me (internal):\n %+v\n", me)

	xmlMe, _ := xml.Marshal(me) // HL
	fmt.Printf("About me (xml):\n %s\n", xmlMe)
}
