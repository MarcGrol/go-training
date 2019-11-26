package main

import "log"

type DB struct {
	mychannel chan string
}

func dowork(db *DB) {
	log.Printf("Perform work")
}

// START OMIT
func New() (*DB, func(), error) {
	db := &DB{
		mychannel: make(chan string),
	}
	cleanupFunc := func() { // HL
		log.Printf("Perform cleanup") // HL
		close(db.mychannel)           // uses variable outside its own body // HL
	} // HL
	return db, cleanupFunc, nil
}

func main() {
	db, cleanup, err := New()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup() // can still reference db.mychannel // HL

	dowork(db)
}

// END OMIT
