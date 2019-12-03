package httpclient

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchPatientSuccess(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/api/patient/123", r.URL.RequestURI())

		returnSuccesResponse(t, w, r)
	}))
	defer testServer.Close()

	// perform request
	client := &PatientClient{HostName:testServer.URL}
	resp, err := client.Fetch("123")

	// validate output as return from fake server
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Marc", resp.Name)
	assert.Equal(t, 42, resp.Age)
}

func returnSuccesResponse(t *testing.T, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := Patient{
		Name: "Marc",
		Age:    42,
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		t.Fatalf("Error encoding response")
	}
}

func TestFetchPatientError(t *testing.T) {
	// create server simulator
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer testServer.Close()

	// perform request
	client := &PatientClient{HostName:testServer.URL}
	_, err := client.Fetch("123")

	// validate output as return from fake server
	assert.Error(t, err)
}
