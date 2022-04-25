package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeysFromMap(t *testing.T) {
	hobbiesByID := getHobbiesByID()
	assert.Equal(t, getHobbiesByIDKeys(hobbiesByID), keysFromMap(hobbiesByID))

	hobbiesOfPeople := getHobbiesOfPeople()
	assert.Equal(t, getHobbiesOfPeopleKeys(hobbiesOfPeople), keysFromMap(hobbiesOfPeople))
}