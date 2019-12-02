package nontestable

import (
	"github.com/google/uuid" // third-party package
	"io/ioutil"
	"time"
)

func Write() error{
	u := uuid.New()
	ft := time.Now().Format(time.RFC3339)

	return ioutil.WriteFile(u.String() + ".txt", []byte(ft), 0644)
}
