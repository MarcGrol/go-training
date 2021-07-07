package testable

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/uuid"
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

func newFiletimeWriter(uuidGenerator UuidGenerator, nower Nower) *filetimeWriter {
	return &filetimeWriter{
		uuidGenerator: uuidGenerator,
		nower:         nower,
	}
}

type realNower struct{}

func (rn *realNower) Now() time.Time {
	return time.Now()
}

type realUuidGenerator struct{}

func (r *realUuidGenerator) Generate() string {
	return uuid.New().String()
}

func (w filetimeWriter) write() (string, error) {
	u := w.uuidGenerator.Generate()
	uppercaseFilename := strings.ToUpper(u) + ".txt"

	now := w.nower.Now()
	futureDate := now.AddDate(1, 2, 3).Format(time.RFC3339)

	return uppercaseFilename, ioutil.WriteFile(uppercaseFilename, []byte(futureDate), 0644)
}

func Write() error {
	// Creates filetimeWriter with default behaviour
	_, err := newFiletimeWriter(&realUuidGenerator{}, &realNower{}).write()
	return err
}
