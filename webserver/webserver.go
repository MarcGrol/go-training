package main

import (
	"fmt"
	"net/http"
	"time"
)

type currentTimeFunc func() time.Time

func getCurrentTime() time.Time {
	return time.Now()
}

func NewTimeHandler(format string) *timeHandler {
	th := new(timeHandler)
	th.nowFunc = getCurrentTime
	return th
}

// START OMIT
type timeHandler struct {
	format  string
	nowFunc currentTimeFunc
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := th.nowFunc().Format(th.format)
	fmt.Fprintf(w, "The time is: "+tm)
}

func main() {
	mux := http.NewServeMux()

	th1123 := NewTimeHandler(time.RFC1123)
	mux.Handle("/time/rfc1123", th1123)

	th3339 := NewTimeHandler(time.RFC3339)
	mux.Handle("/time/rfc3339", th3339)

	http.ListenAndServe(":3000", mux)
}

// END OMIT
