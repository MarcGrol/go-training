package main

import (
	"fmt"
	"strings"
)

func main()  {
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

	peopleOnHobby := map[string][]string{}
	for name, hobbies := range hobbiesOfPeople {
		for _, hobby := range hobbies {
			persons, found := peopleOnHobby[hobby]
			if !found {
				persons = []string{}
			}
			persons = append(persons, name)
			peopleOnHobby[strings.ToLower(hobby)] = persons
		}
	}
	fmt.Println("%+v", peopleOnHobby)
}
