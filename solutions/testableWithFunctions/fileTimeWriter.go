package testable

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/uuid"
)

type filetimeWriter struct {
	uuidGenerator func() string
	nower         func() time.Time
}

func newFiletimeWriter() filetimeWriter {
	// Creates filetimeWriter with default behaviour
	return filetimeWriter{
		uuidGenerator: func() string { return uuid.New().String() },
		nower:         func() time.Time { return time.Now() },
	}
}

func (w filetimeWriter) write() (string, error) {
	u := w.uuidGenerator()
	uppercaseFilename := strings.ToUpper(u) + ".txt"

	now := w.nower()
	futureDate := now.AddDate(1, 2, 3).Format(time.RFC3339)

	return uppercaseFilename, ioutil.WriteFile(uppercaseFilename, []byte(futureDate), 0644)
}

func Write() error {
	_, err := newFiletimeWriter().write()
	return err
}
