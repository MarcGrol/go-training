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

func firstThatStartWith(in []string, prefix string) (foundValue string, exists bool) {
	for _, val := range in {
		if strings.HasPrefix(val, prefix) {
			foundValue = val
			exists = true
			return
		}
	}

	return
}

type hobby string
type name string

func groupBy(hobbiesOfPeople map[hobby][]name) map[string][]string {
	peopleOnHobby := map[string][]string{}
	for name, hobbies := range hobbiesOfPeople {
		for _, hobby := range hobbies {

			hobby := strings.ToLower(hobby)

			//persons, found := peopleOnHobby[hobby]
			//if !found {
			//	persons = []string{}
			//}

			persons, _ := peopleOnHobby[hobby]
			persons = append(persons, name)
			peopleOnHobby[hobby] = persons
		}
	}
	return peopleOnHobby
}
