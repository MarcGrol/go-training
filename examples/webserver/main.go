package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// START OMIT

func main() {
	var router *mux.Router = mux.NewRouter()

	uider := NewBasicUuider()
	patientStore := newPatientStore(uider)
	webService := NewPatientWebService(patientStore)
	webService.RegisterEndpoint(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// END OMIT
