package testable

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	inputTimestamp    = "2019-02-28T00:00:00Z"
	expectedTimestamp = "2020-05-01T00:00:00Z"

	uid              = "abc123"
	expectedFilename = "ABC123.txt"
)

func fakeGenerate() string {
	return uid
}

func fakeNow() time.Time {
	t, _ := time.Parse(time.RFC3339, inputTimestamp)
	return t
}

func TestWithFakes(t *testing.T) {
	err := write(fakeGenerate, fakeNow)
	assert.NoError(t, err)
	defer func() {
		err = os.Remove(expectedFilename) // cleanup
		assert.NoError(t, err)
	}()

	result, err := ioutil.ReadFile(expectedFilename)
	assert.NoError(t, err)

	assert.Equal(t, expectedTimestamp, string(result))
}
