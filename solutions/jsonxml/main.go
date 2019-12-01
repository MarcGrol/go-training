package main

import (
	"bytes"
	"log"
	"os"

	"github.com/MarcGrol/go-training/solutions/jsonxml/person"
)

const (
	filename = "./person.json"
)

func main() {
	// create and populate the struct
	p := person.Person{
		Name:      "Marc",
		Age:       42,
		Interests: []string{"Running", "Cycling", "Hockey"},
	}

	// log the struct and all its attributes as "extended verbose"-printf notation
	log.Printf("struct:%+v\n", p)

	// print the json to stdout
	{
		p.ToJson(os.Stdout)
	}

	// print the json to file
	{
		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Error opening file %s (%s)", filename, err)
		}
		defer f.Close()
		p.ToJson(f)
	}

	// print json to a byte-buffer
	{
		var buf bytes.Buffer
		p.ToJson(&buf)
		log.Printf("buf:%s", buf.Bytes())
	}

}
