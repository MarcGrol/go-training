package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

func main() {

	uider := func() string {
		return uuid.New().String()
	}
	nower := func() time.Time {
		return time.Now()
	}

	emailSender := NewDummyEmailSender()

	store := newClientStore(nower)
	webService := NewPatientService(uider, store, emailSender)

	router := mux.NewRouter()
	webService.RegisterEndpoint(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":80", nil))
}
