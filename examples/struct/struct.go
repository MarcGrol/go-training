package main

import (
	"fmt"
	"time"
)

// START OMIT
type Student struct { // public
	FullName               string
	AddressLine            string
	BirthDate              time.Time
	Study                  Study
	ProgressYear           int
	expectedGraduationDate time.Time // private
}

type Study struct {
	Name string
}

// END OMIT

func CreateStudent(fullName, adddressLine string, birthdate time.Time, studyName string, progressYear int) Student {
	return Student{
		FullName:    fullName,
		AddressLine: adddressLine,
		BirthDate:   birthdate,
		Study: Study{
			Name: studyName,
		},
		ProgressYear: progressYear,
	}
}
func main() {
	student := Student{ // constructor like
		FullName:    "Freek Grol",
		AddressLine: "...., De Bilt, Nederland",
		BirthDate:   time.Now().Add(-1 * 20 * time.Hour * 365 * 18),
		Study: Study{
			Name: "Geography",
		},
		ProgressYear:           3,
		expectedGraduationDate: time.Now().AddDate(1, 2, 3),
	}
	fmt.Printf("%+v", student) // %+v: convenience debugging
}
