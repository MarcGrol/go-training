package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// START OMIT

func main() {
	var router *mux.Router = mux.NewRouter()

	webService := &patientWebService{}
	webService.RegisterEndpoint(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT
