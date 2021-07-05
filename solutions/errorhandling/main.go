package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	filename       = "./solutions/errorhandling/main.go"
	filenameInCaps = filename + ".txt"
)

func main() {
	err := capitalizeFileContent(filename, filenameInCaps)
	if err != nil {
		log.Fatalf("Error reading file %s: %s", filename, err)
	}
}

func capitalizeFileContent(inFilename, outFilename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Error reading file %s: %s", filename, err)
	}

	inCaps := strings.ToUpper(string(data))
	//inCaps := bytes.ToUpper(data)
	log.Printf("In caps: %s", inCaps)

	err = os.WriteFile(filenameInCaps, []byte(inCaps), 0644)
	if err != nil {
		return fmt.Errorf("Error writing file %s: %s", filenameInCaps, err)
	}

	return nil
}