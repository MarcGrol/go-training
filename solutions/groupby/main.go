package main

import (
	"fmt"
	"strings"
)

func main() {
	hobbiesOfPeople := map[string][]string{
		"Julia":  {"voetbal", "tekenen"},
		"Sophie": {"hockey"},
		"Mila":   {"tekenen"},
		"Emma":   {"volleybal", "turnen"},
		"Tess":   {"hardlopen"},
		"Zoë":    {"kunst", "Voetbal"},
		"Noor":   {"voetbal"},
		"Elin":   {"Hockey"},
		"Sara":   {"voetbal", "turnen"},
		"Yara":   {"tekenen"},
	}
	peopleOnHobby := groupBy(hobbiesOfPeople)

	fmt.Printf("%+v", peopleOnHobby)
}

func groupBy(hobbiesOfPeople map[string][]string) map[string][]string {
	peopleOnHobby := map[string][]string{}
	for name, hobbies := range hobbiesOfPeople {
		for _, hobby := range hobbies {

			hobby := strings.ToLower(hobby)

			persons := peopleOnHobby[hobby]
			persons = append(persons, name)
			peopleOnHobby[hobby] = persons
		}
	}
	return peopleOnHobby
}
