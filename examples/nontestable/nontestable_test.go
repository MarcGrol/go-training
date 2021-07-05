package nontestable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {
	// when
	err := Write()

	// then
	assert.NoError(t, err)
	// More tests?

	// cleanup????
}
