package testable

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/uuid"
)

func Write() error {
	return write(
		func() string {
			return uuid.New().String()
		},
		func() time.Time {
			return time.Now()
		},
	)
}

func write(uuider func() string, nower func() time.Time) error {
	u := uuider()
	uppercaseFilename := strings.ToUpper(u) + ".txt"

	now := nower()
	futureDate := now.AddDate(1, 2, 3).Format(time.RFC3339)

	return ioutil.WriteFile(uppercaseFilename, []byte(futureDate), 0644)
}
