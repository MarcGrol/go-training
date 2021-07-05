package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

// START OMIT

func main() {
	var router *mux.Router = mux.NewRouter()

	uider := func() string {
		return uuid.New().String()
	}
	store := newPatientStore(uider)
	nower := func() time.Time {
		return time.Now()
	}

	webService := NewPatientService(uider, nower, store)
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
