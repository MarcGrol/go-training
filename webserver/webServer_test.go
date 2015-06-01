package main

import (
	"fmt"
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

func TestServer(t *testing.T) {
	// mock response
	recorder := httptest.NewRecorder()

	// create a http request that trigger your server
	req, _ := http.NewRequest("GET", "", nil)

	// call subject of test
	th := timeHandler{format: time.RFC3339, nowFunc: mockNow}
	th.ServeHTTP(recorder, req)

	//  verify response
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "The time is: 1971-02-27T14:31:59Z", string(recorder.Body.Bytes()))
	t.Log(fmt.Sprintf("%+v", recorder))

}
