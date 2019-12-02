package testable

import (
	"io/ioutil"
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
	nower Nower
}

func New(uuidGenerator UuidGenerator, nower Nower ) *filetimeWriter {
	return &filetimeWriter{
		uuidGenerator: uuidGenerator,
		nower:nower,
	}
}

func (w filetimeWriter)Write() error {
	u := w.uuidGenerator.Generate()
	ft := w.nower.Now().Format(time.RFC3339)

	return ioutil.WriteFile(u + ".txt", []byte(ft), 0644)
}

func (w filetimeWriter)getFilenameAndContent() error {
	u := w.uuidGenerator.Generate()
	ft := w.nower.Now().Format(time.RFC3339)

	return ioutil.WriteFile(u + ".txt", []byte(ft), 0644)
}