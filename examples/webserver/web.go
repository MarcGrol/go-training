package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func (s *patientWebService) registerTestEndpoint() *mux.Router {
	return s.RegisterEndpoint(mux.NewRouter())
}

// START OMIT

func (s *patientWebService) RegisterEndpoint(router *mux.Router) *mux.Router{
	subRouter := router.PathPrefix("/api/patient").Subrouter()
	subRouter.HandleFunc("/{patientUid}",s.getPatient() ).Methods("GET") // HL
	subRouter.HandleFunc("", s.postPatient()).Methods("POST")
	subRouter.HandleFunc("/{patientUid}", s.putPatient()).Methods("PUT")
	subRouter.HandleFunc("/{patientUid}",s.deletePatient()).Methods("DELETE")
	return router
}

func (s *patientWebService)getPatient() http.HandlerFunc { // HL
	return func(w http.ResponseWriter, r *http.Request) {
		patientUid := mux.Vars(r)["patientUid"] // extract path param
		response, err := s.getPatientOnUID(patientUid) // call business logic // HL
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")  // write response
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
// END OMIT

func (s *patientWebService)postPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse request body
		patient := Patient{}
		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		log.Printf("Patient: %+v", patient)
		// call business logic
		response, err := s.createPatient(patient)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		// write response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (s *patientWebService)putPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	// TODO
	}
}

func (s *patientWebService)deletePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
	}
}
