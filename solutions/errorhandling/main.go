package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const filename = "./solutions/errorhandling/service.go"
	const filenameInCaps = filename + ".txt"
	err := capitalizeFileContent(filename, filenameInCaps)
	if err != nil {
		log.Fatalf("Error reading file %s: %s", filename, err)
	}
}

func capitalizeFileContent(inFilename, outFilename string) error {
	data, err := os.ReadFile(inFilename)
	if err != nil {
		return fmt.Errorf("error reading file %s: %s", inFilename, err)
	}

	inCaps := strings.ToUpper(string(data))
	log.Printf("In caps: %s", inCaps)

	err = os.WriteFile(outFilename, []byte(inCaps), 0644)
	if err != nil {
		return fmt.Errorf("error writing file %s: %s", outFilename, err)
	}

	return nil
}
