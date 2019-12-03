package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	// setup
	router := mux.NewRouter()
	sut := &patientWebService{}
	sut.RegisterEndpoint(router)

	// given
	req, _ := http.NewRequest("GET", "/api/patient/123", nil)

	// when
	recordedResponse := httptest.NewRecorder() // records what was send back by the server
	router.ServeHTTP(recordedResponse, req)

	// then
	assert.Equal(t, http.StatusOK, recordedResponse.Code)
	assert.Contains(t, string(recordedResponse.Body.Bytes()),`{"UID":"`, )
}

func TestPost(t *testing.T) {
	// setup
	router := mux.NewRouter()
	sut := &patientWebService{}
	sut.RegisterEndpoint(router)

	// given
	req, _ := http.NewRequest("POST", "/api/patient", strings.NewReader(
		`{"FullName":"Marc","AddressLine":"Heemstra","Allergies":["pinda"]}`))

	// when
	recordedResponse := httptest.NewRecorder()
	router.ServeHTTP(recordedResponse, req)

	// then
	assert.Equal(t, http.StatusOK, recordedResponse.Code)
	assert.Contains(t, string(recordedResponse.Body.Bytes()),`{"UID":"`, )
}

func TestPut(t *testing.T) {
	// TODO
}

func TestDelete(t *testing.T) {
	// TODO
}