package main

import (
	"time"
)

type Nower interface {
	Now() time.Time
}

func NewNower() Nower {
	return &nower{}
}

type nower struct{}

func (n nower) Now() time.Time {
	return time.Now()
}
