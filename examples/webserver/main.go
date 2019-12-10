package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// START OMIT

func main() {
	var router *mux.Router = mux.NewRouter()

	uider := NewBasicUuider()
	patientStore := newPatientStore(uider)
	webService := NewPatientWebService(patientStore)
	webService.RegisterEndpoint(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT
