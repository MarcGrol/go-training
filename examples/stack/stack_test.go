package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := Stack{}
	s.Push("world!")
	assert.Equal(t, 1, s.Size())

	s.Push("Hello, ")
	assert.Equal(t, 2, s.Size())

	hello := s.Pop()
	assert.Equal(t, "Hello, ", hello)
	assert.Equal(t, 1, s.Size())

	world := s.Pop()
	assert.Equal(t, "world!", world)
	assert.Equal(t, 0, s.Size())

	t.Logf("stack:%+v", s)
}
