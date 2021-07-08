package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"encoding/json"
	"encoding/xml"
)

type Person struct {
	Name            string    `json:"name"      xml:"PersonName"`
	BirthDate       MyDate    `json:"birth-date" xml:"Birthdate"`                      // MyDate override standard d xml json serializers
	NextAppointment time.Time `json:"next-appointment" xml:"NextAppointmentTimestamp"` // RFC3339 (2006-01-02T15:04:05Z) works out oof the box
	Interests       []string  `json:"interests" xml:"PersonInterests"`
	Children        []Child   `json:"children"  xml:"Person_Children"`
}

type Child struct {
	Name string `json:"name"          xml:"ChildName"`
	Age  int    `json:"age,omitempty" xml:"Child_Age,omitempty"`
}

func main() {
	var jsonMe string = `
	{
		"name":"Marc Grol",
        "birth-date": "27-02-1971",
		"next-appointment":"2006-01-02T15:04:05Z",
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
		log.Fatalf("Unmarshall error:%s", err)
	}

	dumpAsJson(me)

	xmlMe, err := xml.MarshalIndent(me, "", "\t") // HL
	if err != nil {
		log.Fatalf("Marshall error:%s", err)
	}
	fmt.Printf("About me (xml):\n %s\n", xmlMe)
}

func dumpAsJson(p Person) {
	json.NewEncoder(os.Stdout).Encode(p)
}
