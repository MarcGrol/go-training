package main

import (
	"bytes"
	"log"
	"os"
)

const (
	filename       = "./solutions/errorhandling/main.go"
	filenameInCaps = filename + ".txt"
)

func main() {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file %s: %s", filename, err)
	}

	inCaps := bytes.ToUpper(data)
	log.Printf("In caps: %s", inCaps)

	err = os.WriteFile(filenameInCaps, inCaps, 0644)
	if err != nil {
		log.Fatalf("Error writing file %s: %s", filenameInCaps, err)
	}

}
