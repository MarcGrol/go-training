package directoryplitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectorySplitter(t *testing.T) {
	randomString := "a/b/c"
	splitString := DirectorySplitter(randomString)

	assert.Equal(t, []string{"a", "b", "c"}, splitString)

}
