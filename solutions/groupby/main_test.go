package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy(t *testing.T) {
	hobbiesOfPeople := map[string][]string{
		"Julia":  {"Voetbal", "tekenen"},
		"Sophie": {"voetbal", "turnen"},
	}
	peopleOnHobby := groupBy(hobbiesOfPeople)

	assert.Len(t, peopleOnHobby["tekenen"], 1, "tekenen")
	assert.Len(t, peopleOnHobby["turnen"], 1, "turnen")
	assert.Len(t, peopleOnHobby["voetbal"], 2, "voetbal")
	assert.Len(t, peopleOnHobby["hockey"], 0, "hockey")
}
