package main

import (
	"fmt"
	"sort"
)

func main() {
	hobbiesOfPeople := getHobbiesOfPeople()
	hobbiesByID := getHobbiesByID()

	fmt.Println(getHobbiesOfPeopleKeys(hobbiesOfPeople))
	fmt.Println(getHobbiesByIDKeys(hobbiesByID))

}

func getHobbiesOfPeople() map[string][]string {
	return map[string][]string{
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
}
func getHobbiesOfPeopleKeys(hobbiesOfPeople map[string][]string) []string {
	var keys []string
	for key := range hobbiesOfPeople {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func getHobbiesByID() map[int]string {
	return map[int]string{
		34: "voetbal",
		2:  "hockey",
		43: "tekenen",
		21: "volleybal",
		35: "turnen",
		24: "kunst",
		93: "hardlopen",
	}
}
func getHobbiesByIDKeys(hobbiesByID map[int]string) []int {
	var keys []int
	for key := range hobbiesByID {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	return keys
}
