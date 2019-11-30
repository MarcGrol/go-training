package benchmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	bs = BigStruct{}
)

// Regular tests
func TestDoCalculationByValue(t *testing.T) {
	value := bs.DoCalculationByValue()
	assert.Equal(t, 1984, value)
}

func TestDoCalculationByReference(t *testing.T) {
	value := (&bs).DoCalculationByReference()
	assert.Equal(t, 42, value)
}

// Naming convention: starts with "Benchmark" and has "b *testing.B" as parameter
// trigger benchmark with: go test -bench=.
func BenchmarkDoCalculationByValue(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		bs.DoCalculationByValue()
	}
}

func BenchmarkDoCalculationByReference(b *testing.B) {
	for n := 0; n < b.N; n++ {
		(&bs).DoCalculationByReference()
	}
}
