package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	// mock response
	recorder := httptest.NewRecorder()

	// create a http request that trigger your server
	req, _ := http.NewRequest("GET", "/api/patient/123", nil)

	// call subject of test
	webservice := &patientWebService{}
	webservice.registerTestEndpoint().ServeHTTP(recorder, req)

	//  verify response
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, string(recorder.Body.Bytes()),`{"UID":"`, )
}

func TestPost(t *testing.T) {
	// mock response
	recorder := httptest.NewRecorder()

	// create a http request that trigger your server
	req, _ := http.NewRequest("POST", "/api/patient", strings.NewReader(
		`{"FullName":"Marc","AddressLine":"Heemstra","Allergies":["pinda"]}`))

	// call subject of test
	webservice := &patientWebService{}
	webservice.registerTestEndpoint().ServeHTTP(recorder, req)

	//  verify response
	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, string(recorder.Body.Bytes()),`{"UID":"`, )
}

func TestPut(t *testing.T) {
	// TODO
}

func TestDelete(t *testing.T) {
	// TODO
}