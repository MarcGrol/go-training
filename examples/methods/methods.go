package main

import (
	"log"
	"time"
)

// START OMIT
type Patient struct {
	Name        string
	YearBorn    int
	IsHealthy   bool
	LastChecked time.Time
}

func (p Patient) HasHighRiskOnDisease() bool { // no side effect
	return (time.Now().Year() - p.YearBorn) > 70
}

func (p *Patient) MarkHealthy() { // has side effect
	p.IsHealthy = true
	p.LastChecked = time.Now()
}

func main() {
	opa := Patient{
		Name:     "Hans",
		YearBorn: 1940,
	}
	log.Printf("high-risk: %+v\n", opa.HasHighRiskOnDisease())
	opa.MarkHealthy()
	log.Printf("after: %+v\n", opa)
}

// END OMIT
