package packaging

import (
	"fmt"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "The time is: "+tm)
}

func serveTimeRequests() {
	mux := http.NewServeMux()
	mux.Handle("/time", http.HandlerFunc(timeHandler))
	http.ListenAndServe(":3000", mux)
}
