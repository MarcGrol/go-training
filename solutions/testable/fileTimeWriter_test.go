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

type mockNower struct{}

func (m mockNower) Now() time.Time {
	t, _ := time.Parse(time.RFC3339, inputTimestamp)
	return t
}

type mockUider struct{}

func (u mockUider) Generate() string {
	return uid
}

func TestWithHandmadeStubs(t *testing.T) {
	ftw := newFiletimeWriter(&mockUider{}, &mockNower{})

	filename, err := ftw.write()
	assert.NoError(t, err)
	defer func() {
		err = os.Remove(expectedFilename) // cleanup
		assert.NoError(t, err)
	}()

	assert.Equal(t, expectedFilename, filename)
	result, err := ioutil.ReadFile(expectedFilename)
	assert.NoError(t, err)

	assert.Equal(t, expectedTimestamp, string(result))
}
