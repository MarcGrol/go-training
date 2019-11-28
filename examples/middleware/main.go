package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

// START OMIT
func main() {
	svc := applicativeService{}
	chain := alice.New(measuringFilter, authenticationFilter).Then(svc) // HL

	http.Handle("/", chain)
	http.ListenAndServe(":8080", nil)
}

func measuringFilter(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Start measuring")
		defer log.Printf("Done measuring")

		t1 := time.Now()
		next.ServeHTTP(w, r) // HL
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

// END OMIT

type applicativeService struct{}

func (bl applicativeService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Enter business logic")
	defer log.Printf("Leave business logic")

	// Extract headers
	role := r.Header.Get("X-enduser-role")
	userUid := r.Header.Get("X-enduser-uid")

	// Call business logic
	response := bl.doit(role, userUid)

	// return response
	fmt.Fprintf(w, response)
}

func (bl applicativeService) doit(role, userUID string) string {
	response := fmt.Sprintf("Welkom %s %s", role, userUID)
	return response
}

func authenticationFilter(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Start authenthication")
		defer log.Printf("Done authenthication")
		if isProtected(r.RequestURI) {
			token := extractToken(r)
			if !isTokenValid(token) {
				w.WriteHeader(http.StatusForbidden)
				return
			}
			userUID, role := extractUserInfoFromToken(token)
			r.Header.Set("X-enduser-uid", userUID)
			r.Header.Set("X-enduser-role", role)
			log.Printf("User %s with role %s is authenticated", userUID, role)
		}
		next.ServeHTTP(w, r) // HL
	}
	return http.HandlerFunc(fn)
}

func isProtected(url string) bool {
	return true
}

func extractToken(r *http.Request) string {
	return "123"
}

func isTokenValid(token string) bool {
	return true
}

func extractUserInfoFromToken(token string) (string, string) {
	return "1234321", "patient"
}
