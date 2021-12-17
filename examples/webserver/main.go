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
	router := mux.NewRouter()

	uider := func() string {
		return uuid.New().String()
	}
	nower := func() time.Time {
		return time.Now()
	}

	store := newPatientStore(nower)
	webService := NewPatientService(uider, store)
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
