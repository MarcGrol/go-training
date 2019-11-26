package main

import "log"

// START OMIT

func enter(name string) string {
	log.Printf("enter %s", name)
	return name
}

func leave(name string) {
	log.Printf("leave %s", name)
}

func main() {
	defer leave(enter("main")) // HL
	log.Printf("in main")
}

// END OMIT
