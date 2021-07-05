package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type predicatbleUider struct{}

func (u predicatbleUider) Create() string {
	return "1"
}

func testUUIDer() string {
	return "123"
}

func testNower() time.Time {
	location, _ := time.LoadLocation("Europe/Amsterdam")
	t, _ := time.ParseInLocation("2006-01-02T15:04:05", "2016-02-27T00:00:00", location)
	return t
}

func TestGet(t *testing.T) {
	// setup
	router := mux.NewRouter()
	patientStore := newPatientStore(predicatbleUider{})
	sut := NewPatientService(testUUIDer, testNower, patientStore)
	sut.RegisterEndpoint(router)

	// given
	_, err := patientStore.Put(context.TODO(), Patient{
		FullName:    "Me",
		AddressLine: "Here",
		Allergies:   []string{"trouble"},
	})
	assert.NoError(t, err)

	// when
	req, err := http.NewRequest("GET", "/api/patient/1", nil)
	assert.NoError(t, err)
	recordedResponse := httptest.NewRecorder() // records what was send back by the server
	router.ServeHTTP(recordedResponse, req)

	// then
	assert.Equal(t, http.StatusOK, recordedResponse.Code)
	assert.Contains(t, string(recordedResponse.Body.Bytes()), `{"uid":"`)
}

func TestPost(t *testing.T) {
	// setup
	router := mux.NewRouter()
	patientStore := newPatientStore(predicatbleUider{})
	sut := NewPatientService(testUUIDer, testNower, patientStore)
	sut.RegisterEndpoint(router)

	// given
	err := patientStore.Remove(context.TODO(), "1")
	assert.NoError(t, err)

	// when
	req, err := http.NewRequest("POST", "/api/patient", strings.NewReader(
		`{"FullName":"Marc","AddressLine":"Weezenhof","Allergies":["gedoe"]}`))
	assert.NoError(t, err)
	recordedResponse := httptest.NewRecorder()
	router.ServeHTTP(recordedResponse, req)

	// then
	assert.Equal(t, http.StatusOK, recordedResponse.Code)
	assert.Contains(t, string(recordedResponse.Body.Bytes()), `"Weezenhof"`)
}

func TestPut(t *testing.T) {
	// setup
	router := mux.NewRouter()
	patientStore := newPatientStore(predicatbleUider{})
	sut := NewPatientService(testUUIDer, testNower, patientStore)
	sut.RegisterEndpoint(router)

	// given
	_, err := patientStore.Put(context.TODO(), Patient{
		FullName:    "Me",
		AddressLine: "Here",
		Allergies:   []string{"trouble"},
	})
	assert.NoError(t, err)

	// when
	req, err := http.NewRequest("PUT", "/api/patient/1", strings.NewReader(
		`{"FullName":"Marc","AddressLine":"Heemstra","Allergies":["pinda"]}`))
	assert.NoError(t, err)
	recordedResponse := httptest.NewRecorder()
	router.ServeHTTP(recordedResponse, req)

	// then
	assert.Equal(t, http.StatusOK, recordedResponse.Code)
	assert.Contains(t, string(recordedResponse.Body.Bytes()), `"Heemstra"`)
}

func TestDelete(t *testing.T) {
	// setup
	router := mux.NewRouter()
	patientStore := newPatientStore(predicatbleUider{})
	sut := NewPatientService(testUUIDer, testNower, patientStore)
	sut.RegisterEndpoint(router)

	// given
	_, err := patientStore.Put(context.TODO(), Patient{
		FullName:    "Me",
		AddressLine: "Here",
		Allergies:   []string{"trouble"},
	})
	assert.NoError(t, err)

	// when
	req, err := http.NewRequest("DELETE", "/api/patient/1", nil)
	assert.NoError(t, err)
	recordedResponse := httptest.NewRecorder() // records what was send back by the server
	router.ServeHTTP(recordedResponse, req)

	// then
	assert.Equal(t, http.StatusOK, recordedResponse.Code)
	_, found, err := patientStore.GetOnUid(context.TODO(), "1")
	assert.NoError(t, err)
	assert.False(t, found)
}
