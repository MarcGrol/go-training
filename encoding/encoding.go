// hello.go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"      xml:"PersonName"`
	Interests []string `json:"interests" xml:"PersonInterests"`
	Children  []Child  `json:"children"  xml:"Person_Children"`
}

type Child struct {
	Name string `json:"name"          xml:"ChildName"`
	Age  int    `json:"age,omitempty" xml:"Child_Age,omitempty"`
}

func main() {
	var jsonMe string = `
	{
		"name":"Marc Grol",
	    "interests":["Running","Golang"],
	    "children":[
	    	{"name":"Pien","shirtNumber":12,"age":5},
	    	{"name":"Tijl","shirtNumber":9},
	    	{"name":"Freek"}
	    ]
	}`

	var me Person
	_ = json.Unmarshal([]byte(jsonMe), &me) // HL

	fmt.Printf("About me (internal):\n %+v\n", me)

	xmlMe, _ := xml.MarshalIndent(me, "", "\t") // HL
	fmt.Printf("About me (xml):\n %s\n", xmlMe)
}
