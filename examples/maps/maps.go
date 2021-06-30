package main

import "fmt"

// START OMIT
func main() {
	// initialize map // HL
	studentsOnSchool := map[string][]string{
		"Jordan": []string{"Jelle", "Hidde"},
		"HNL":    []string{"Tijl"},
	}
	fmt.Printf("1: %+v\n", studentsOnSchool) // %+v debugging convenience

	// add map entry // HL
	studentsOnSchool["OLZ"] = []string{"Pien"}

	// remove map entry  // HL
	delete(studentsOnSchool, "HNL")

	// get map entry  // HL
	jordanStudents, found := studentsOnSchool["Jordan"]
	if !found {
		jordanStudents = []string{}
	}

	// iterate map // HL
	for key, value := range studentsOnSchool {
		fmt.Printf("4: %s - %v\n", key, value)
	}
}

// END OMIT
