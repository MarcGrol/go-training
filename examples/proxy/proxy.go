package main

import (
	"net/http"
	"time"
)

// START OMIT
func main() {
	statsCollector := NewStatsCollector()
	measuredHandler := MeasuredHttpProcessor(HandleHttpRequest, statsCollector) // HL
	http.HandleFunc("/api", measuredHandler)                                    // Only accepts standard http.HandlerFunc
}

func MeasuredHttpProcessor(httpFunc func(w http.ResponseWriter, r *http.Request), sc *StatsCollector) func(w http.ResponseWriter, r *http.Request) { // HL
	return func(w http.ResponseWriter, r *http.Request) { // return a standard http.HandlerFunc
		before := time.Now()
		httpFunc(w, r)
		after := time.Now()
		sc.Collect(r.URL.Path, after.Sub(before)) // Upload duration to monitoring sub-system
	}
}

func HandleHttpRequest(w http.ResponseWriter, r *http.Request) { // HL
	// TODO Read and decode http request
	// TODO Delegate to business logic
	// TODO Encode and write response
}

// END OMIT

type StatsCollector struct{}

func NewStatsCollector() *StatsCollector {
	return &StatsCollector{}
}

func (sc *StatsCollector) Collect(path string, dur time.Duration) {

}
