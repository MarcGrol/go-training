package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *clientWebService) RegisterEndpoint(router *mux.Router) *mux.Router {
	subRouter := router.PathPrefix("/api/client").Subrouter()
	subRouter.HandleFunc("/{clientUid}", s.getPatient()).Methods("GET") // HL
	subRouter.HandleFunc("", s.postPatient()).Methods("POST")
	subRouter.HandleFunc("/{clientUid}", s.putPatient()).Methods("PUT")
	subRouter.HandleFunc("/{clientUid}", s.deletePatient()).Methods("DELETE")
	return router
}

func (s *clientWebService) getPatient() http.HandlerFunc { // HL
	c := context.Background()
	return func(w http.ResponseWriter, r *http.Request) {
		clientUid := mux.Vars(r)["clientUid"]                   // extract path param
		response, found, err := s.getPatientOnUID(c, clientUid) // call business logic // HL
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json") // write response
		json.NewEncoder(w).Encode(response)
	}
}

func (s *clientWebService) postPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := context.Background()
		// parse request body
		client := Client{}
		err := json.NewDecoder(r.Body).Decode(&client)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// call business logic
		response, err := s.createClient(c, client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// write response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func (s *clientWebService) putPatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := context.Background()

		// parse request body
		client := Client{}
		err := json.NewDecoder(r.Body).Decode(&client)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// parse url
		clientUid := mux.Vars(r)["clientUid"] // extract path param
		client.UID = clientUid

		// call business logic
		err = s.modifyClientOnUid(c, client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// write response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(client)
	}
}

func (s *clientWebService) deletePatient() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := context.Background()
		clientUid := mux.Vars(r)["clientUid"] // extract path param

		// call business logic
		err := s.deleteClientOnUid(c, clientUid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
