package main

import (
	"expvar"
	"io"
	"net/http"
)

var (
	hits               = expvar.NewInt("hits")
	lastCaller         = expvar.NewString("last_user")
	myCombinedCounters = expvar.NewMap("my_counters")
)

func handleHit(w http.ResponseWriter, r *http.Request) {
	lastCaller.Set(r.RemoteAddr)
	hits.Add(1)
	myCombinedCounters.Add("my_hits", 1)
	io.WriteString(w, hits.String())
}

func init() {
	//hits.Add(42)
	lastCaller.Set("unknown")
	myCombinedCounters.Add("my_hits", 42)
}

func main() {

	http.HandleFunc("/doit", handleHit)
	http.ListenAndServe(":8000", nil)

	// view expvars on: http://localhost:8000/debug/vars

}
