package testable

import (
	"io/ioutil"
	"strings"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -source=fileTimeWriter.go -destination=fileTimeWriterMocks.go -package=testable UuidGenerator,Nower

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
	ftw := filetimeWriter{
		uuidGenerator: &realUuidGenerator{},
		nower:         &realNower{},
	}
	_, err := ftw.write()
	return err
}
