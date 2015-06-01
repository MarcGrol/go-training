package main

import (
	"fmt"
	"net/http"
	"time"
)

// provide a hook for time that can be overridden in tests
var nowFunc func() time.Time = getCurrentTime

// default implementation
func getCurrentTime() time.Time {
	return time.Now()
}

// START OMIT
type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := nowFunc().Format(th.format)
	fmt.Fprintf(w, "The time is: "+tm)
}

func main() {
	mux := http.NewServeMux()

	th1123 := &timeHandler{format: time.RFC1123}
	mux.Handle("/time/rfc1123", th1123)

	th3339 := &timeHandler{format: time.RFC3339}
	mux.Handle("/time/rfc3339", th3339)

	http.ListenAndServe(":3000", mux)
}

// END OMIT
