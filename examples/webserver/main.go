package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// START OMIT

func main() {
	webService := &patientWebService{}

	var router = mux.NewRouter()
	webService.RegisterEndpoint(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT
