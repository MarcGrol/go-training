package slowapi

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Sum can be used to simulate a slow action
func Sum(a, b int) int {
	// Sleep-time:
	// - min: 500 msec
	// - average 1000 msec
	// - max: 1500 msec
	sleepDurationInMillisec := 500 + (rand.Intn(1000))
	time.Sleep(time.Duration(sleepDurationInMillisec) * time.Millisecond)
	return a + b
}
