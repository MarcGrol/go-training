package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	testCases := []struct {
		input  string
		status int
		output string
	}{
		{input: "/dylan@nexuzhealth.be", status: http.StatusOK, output: "Hallo dylan van nexushealth"},
		{input: "/mgrol@xebia.com", status: http.StatusOK, output: "Hi there\n"},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.input), func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tc.input, nil)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			myhandler(resp, req)

			assert.Equal(t, tc.status, resp.Code)
			assert.Equal(t, tc.output, string(resp.Body.Bytes()))
		})
	}
}

func BenchmarkMyHandler(b *testing.B) {
	req, _ := http.NewRequest(http.MethodGet, "/dylan@nexushealth.be", nil)
	for n := 0; n < b.N; n++ {
		resp := httptest.NewRecorder()
		myhandler(resp, req)
	}
}
