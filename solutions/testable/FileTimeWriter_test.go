package testable

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	expectedTimestamp = "2019-02-27T00:00:00Z"
)

type mockNower struct{}

func (m mockNower) Now() time.Time {
	t, _ := time.Parse(time.RFC3339, expectedTimestamp)
	return t
}

type mockUider struct{}

func (u mockUider) Generate() string {
	return "1"
}

func TestFileWritten(t *testing.T) {
	ftw := New(&mockUider{}, &mockNower{})

	err := ftw.Write()
	assert.NoError(t, err)
	defer func() {
		err = os.Remove("1.txt") // cleanup
		assert.NoError(t, err)
	}()

	result, err := ioutil.ReadFile("1.txt")
	assert.NoError(t, err)

	assert.Equal(t, expectedTimestamp, string(result))
}
