package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// START OMIT

func (s *patientWebService) RegisterEndpoint(router *mux.Router) *mux.Router {
	subRouter := router.PathPrefix("/api/patient").Subrouter()
	subRouter.HandleFunc("/{patientUid}", s.getPatient()).Methods("GET") // HL
	subRouter.HandleFunc("", s.postPatient()).Methods("POST")
	subRouter.HandleFunc("/{patientUid}", s.putPatient()).Methods("PUT")
	subRouter.HandleFunc("/{patientUid}", s.deletePatient()).Methods("DELETE")
	return router
}

func (s *patientWebService) getPatient() http.HandlerFunc { // HL
	c := context.Background()
	return func(w http.ResponseWriter, r *http.Request) {
		patientUid := mux.Vars(r)["patientUid"]                  // extract path param
		response, found, err := s.getPatientOnUID(c, patientUid) // call business logic // HL
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json") // write response
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

// END OMIT

func (s *patientWebService) postPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := context.Background()
		// parse request body
		patient := Patient{}
		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// call business logic
		response, err := s.createPatient(c, patient)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// write response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *patientWebService) putPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := context.Background()

		// parse request body
		patient := Patient{}
		err := json.NewDecoder(r.Body).Decode(&patient)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// parse url
		patientUid := mux.Vars(r)["patientUid"] // extract path param
		patient.UID = patientUid

		// call business logic
		err = s.modifyPatientOnUid(c, patient)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// write response
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(patient)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *patientWebService) deletePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := context.Background()
		patientUid := mux.Vars(r)["patientUid"] // extract path param

		// call business logic
		err := s.deletePatientOnUid(c, patientUid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
