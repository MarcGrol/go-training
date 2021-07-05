package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START OMIT
func SimulateSlowAction(a, b int) int {
	sleepDurationInMillisec := 500 + (rand.Intn(1000))
	time.Sleep(time.Duration(sleepDurationInMillisec) * time.Millisecond) // sleep approx. 1 sec
	return a * b
}

// END OMIT
