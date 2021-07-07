package testable

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

const (
	inputTimestamp    = "2019-02-28T00:00:00Z"
	expectedTimestamp = "2020-05-01T00:00:00Z"

	uid              = "abc123"
	expectedFilename = "ABC123.txt"
)

func TestWithGeneratedMocks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uuider := NewMockUuidGenerator(ctrl)
	nower := NewMockNower(ctrl)

	// setup expectations
	uuider.EXPECT().Generate().Return(uid)
	nower.EXPECT().Now().Return(func() time.Time {
		t, _ := time.Parse(time.RFC3339, inputTimestamp)
		return t
	}())

	sut := filetimeWriter{
		uuidGenerator: uuider,
		nower:         nower,
	}

	filename, err := sut.write()
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
