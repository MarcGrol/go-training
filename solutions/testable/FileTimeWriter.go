package testable

import (
	"io/ioutil"
	"strings"
	"time"
)

type UuidGenerator interface {
	Generate() string
}

type Nower interface {
	Now() time.Time
}

type filetimeWriter struct {
	uuidGenerator UuidGenerator
	nower         Nower
}

func New(uuidGenerator UuidGenerator, nower Nower) *filetimeWriter {
	return &filetimeWriter{
		uuidGenerator: uuidGenerator,
		nower:         nower,
	}
}

func (w filetimeWriter) Write() (string, error) {
	u := w.uuidGenerator.Generate()
	ft := w.nower.Now()

	uppercaseFilename := strings.ToUpper(u) + ".txt"

	futureDate := ft.AddDate(1, 2, 3).Format(time.RFC3339)

	return uppercaseFilename, ioutil.WriteFile(uppercaseFilename, []byte(futureDate), 0644)
}
