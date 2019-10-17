package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
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
	err := json.Unmarshal([]byte(jsonMe), &me) // HL
	if err != nil {
		log.Fatal(err)
	}

	xmlMe, err := xml.MarshalIndent(me, "", "\t") // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("About me (xml):\n %s\n", xmlMe)
}
