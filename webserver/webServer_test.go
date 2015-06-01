package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func mockNow() time.Time {
	t, _ := time.Parse(time.RFC3339, "1971-02-27T14:31:59Z")
	return t
}

func TestServerRFC3339(t *testing.T) {
	// mock response
	recorder := httptest.NewRecorder()

	// create a http request that trigger your server
	req, _ := http.NewRequest("GET", "", nil)

	// call subject of test
	nowFunc = mockNow
	th := timeHandler{format: time.RFC3339}
	th.ServeHTTP(recorder, req)

	//  verify response
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "The time is: 1971-02-27T14:31:59Z", string(recorder.Body.Bytes()))
}

func TestServerRFC1123(t *testing.T) {
	// mock response
	recorder := httptest.NewRecorder()

	// create a http request that trigger your server
	req, _ := http.NewRequest("GET", "", nil)

	// call subject of test
	nowFunc = mockNow
	th := timeHandler{format: time.RFC1123}
	th.ServeHTTP(recorder, req)

	//  verify response
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "The time is: Sat, 27 Feb 1971 14:31:59 UTC", string(recorder.Body.Bytes()))
}
