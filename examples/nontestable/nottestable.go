package nontestable

import (
	"io/ioutil"
	"time"

	"github.com/google/uuid" // third-party package
)

func Write() error {
	u := uuid.New()
	ft := time.Now().Format(time.RFC3339)
	return ioutil.WriteFile(u.String()+".txt", []byte(ft), 0644)
}
