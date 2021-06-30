package main

import "fmt"

func main()  {
	hobbiesOfPeople := map[string][]string{
		"Julia":  {"voetbal", "tekenen"},
		"Sophie": {"hockey"},
		"Mila":   {"tekenen"},
		"Emma":  {"volleybal", "turnen"},
		"Tess":   {"hardlopen"},
		"Zoë":    {"kunst", "voetbal"},
		"Noor":   {"voetbal"},
		"Elin":   {"hockey"},
		"Sara":   {"voetbal", "turnen"},
		"Yara":   {"tekenen"},
	}

	peopleOnHobby := map[string][]string{}
	for name, hobbies := range hobbiesOfPeople {
		for _, hobby := range hobbies {
			peopleOnHobby[hobby] = append(peopleOnHobby[hobby], name)
		}
	}
	fmt.Println("%+v", peopleOnHobby)
}
