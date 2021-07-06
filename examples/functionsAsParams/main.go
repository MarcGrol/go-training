package main

import (
	"log"

	"github.com/MarcGrol/go-training/examples/functionsAsParams/rigidFramework"
)

// START OMIT
func main() {
	db := &Database{}
	userId := "123" // Read from cli-args?

	// Potentially complex business logic is packed as a simple variable // HL
	businessLogicFunc := func() error {
		user, err := db.Get(userId) // Make ues of database and userId in outer scopee
		if err != nil {
			return err
		}
		// TODO Do some smart adjustment of user
		return db.Put(user)
	}

	// Framework only accepts simple function with signature func() error // HL
	err := rigidFramework.Execute(businessLogicFunc)
	if err != nil {
		log.Fatal(err)
	}
}

// END OMIT

type Database struct {
}

func (db *Database) Get(uid string) (string, error) {
	return "", nil
}

func (db *Database) Put(user string) error {
	return nil
}
