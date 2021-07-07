package testable

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFileWritten(t *testing.T) {
	sut := filetimeWriter{
		uuidGenerator: func() string {
			return "abc123"
		}, nower: func() time.Time {
			t, _ := time.Parse(time.RFC3339, "2019-02-28T00:00:00Z")
			return t
		},
	}

	filename, err := sut.write()
	assert.NoError(t, err)
	defer func() {
		err = os.Remove("ABC123.txt") // cleanup
		assert.NoError(t, err)
	}()

	assert.Equal(t, "ABC123.txt", filename)
	result, err := ioutil.ReadFile("ABC123.txt")
	assert.NoError(t, err)

	assert.Equal(t, "2020-05-01T00:00:00Z", string(result))
}
