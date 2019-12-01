package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const (
	filename = "./exercises/errorhandling/main.go"
	filenameInCaps = filename + ".txt"
)

func main() {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file %s: %s", filename, err)
	}

	inCaps := strings.ToUpper(string(data))
	log.Printf("In caps: %s", inCaps)

	err = ioutil.WriteFile(filenameInCaps, []byte(inCaps), 0644)
	if err != nil {
		log.Fatalf("Error writing file %s: %s", filenameInCaps, err)
	}
}

