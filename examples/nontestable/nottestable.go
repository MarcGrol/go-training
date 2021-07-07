package nontestable

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/uuid" // third-party package
)

func Write() error {

	// Determine the current date and add 1 year, 2 month and 3 days
	future := time.Now().AddDate(1, 2, 3)

	// Format date in iso 3339 format
	formattedDate := future.Format(time.RFC3339)

	// Create a file with a random upper-case name
	filename := strings.ToUpper(uuid.New().String())

	// write adjusted formatted date to file with random uppercase name
	return ioutil.WriteFile(filename+".txt", []byte(formattedDate), 0644)
}
