package main

import "fmt"

// START OMIT
func main() {
	studentsOnSchool := map[string][]string{
		"Jordan": []string{"Jelle", "Hidde"},
		"HNL": []string{"Tijl"},
	}
	fmt.Printf("1: %+v\n", studentsOnSchool) // %+v debugging convenience

	studentsOnSchool["OLZ"] = []string{"Pien"} // add map entry // HL
	fmt.Printf("2: %+v\n", studentsOnSchool) // %+v debugging convenience

	delete(studentsOnSchool, "HNL") // remove map entry  // HL
	fmt.Printf("3: %+v\n", studentsOnSchool) // %+v debugging convenience

	jordanStudents, found := studentsOnSchool["Jordan"] // get map entry  // HL
	if !found {
		jordanStudents = []string{}
	}
	jordanStudents = append(jordanStudents, "Koen")
	studentsOnSchool["Jordan"] = jordanStudents // put map entry // HL

	for key, value := range studentsOnSchool { // iterate map // HL
		fmt.Printf("4: %s - %v\n",key, value)
	}
}
// END OMIT
