package nontestable

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/uuid" // third-party package
)

func Write() error {
	filename := strings.ToUpper(uuid.New().String())
	future := time.Now().AddDate(1, 2, 3).Format(time.RFC3339)
	return ioutil.WriteFile(filename+".txt", []byte(future), 0644)
}
