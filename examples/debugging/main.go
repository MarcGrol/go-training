package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", myhandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

var re = regexp.MustCompile("^(.+)@nexuzhealth.be$")

func myhandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path[1:]
	match := re.FindAllStringSubmatch(path, -1)
	if match != nil {
		fmt.Fprintf(w, "Hallo %s van nexuzhealth", match[1])
		return
	}
	fmt.Fprintln(w, "Hi there")
}
