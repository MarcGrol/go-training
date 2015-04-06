package main

import (
	"flag"
	"log"
)

func main() {
	login := flag.String("login", "", "GitHub login of user")
	flag.Parse()

	if login == "" {
		log.Fatal("must specify login")
	}

	log.Println("Looking up GitHub user: ", login)
}
