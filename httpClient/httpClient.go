package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Origin  string    `json:"origin"`
    Method  string    `json:"method"`
	Url     string    `json:"url"`
	Args    Arguments `json:"args"`
	Headers Headers   `json:"headers"`
    Body    string    `json:"body"`
}

type Arguments struct {
	Arg1 string `json:"arg1"`
	Arg2 string `json:"arg2"`
}

type Headers struct {
	Accept         string `json:"Accept"`
	AcceptEncoding string `json:"Accept-Encoding"`
	AcceptLanguage string `json:"Accept-Language"`
	Cookie         string `json:"Cookie"`
	Host           string `json:"Host"`
	UserAgent      string `json:"User-Agent"`
}

func getIt(url string) (*Response, error) {
	// Perform HTTP GET
	httpResponse, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching: %s", err)
		return nil, err
	}

	// Decode json

	dec := json.NewDecoder(httpResponse.Body)
	var resp Response
	err = dec.Decode(&resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func main() {
	resp, err := getIt("http://httpbin.org/get?arg1=1&arg2=2")
	if err != nil {
	} else {
		log.Printf("%+v", resp)
	}
}
