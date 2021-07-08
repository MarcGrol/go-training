package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type MyDate struct {
	UnderlyingTime time.Time
}

func (t MyDate) MarshalJSON() ([]byte, error) {
	bytes, err := json.Marshal(t.UnderlyingTime.Format("02-01-2006"))
	if err != nil {
		return nil, fmt.Errorf("Error encoding my-date: %+v: %s", t, err)
	}
	return bytes, nil
}

func (t *MyDate) UnmarshalJSON(data []byte) error {
	var dateAsString string
	err := json.Unmarshal(data, &dateAsString)
	if err != nil {
		return fmt.Errorf("Error decoding my-date: %+v: %s", data, err)
	}
	dateAsString = strings.Trim(dateAsString, `"`)
	date, err := time.Parse("02-01-2006", dateAsString)
	if err != nil {
		return fmt.Errorf("Error parsing my-date: %s: %s", dateAsString, err)
	}
	*t = MyDate{
		UnderlyingTime: date,
	}
	return nil
}

func (t MyDate) MarshalText() ([]byte, error) {
	formatted := t.UnderlyingTime.Format("02-01-2006")
	return []byte(formatted), nil
}

func (t *MyDate) UnmarshalText(data []byte) error {
	var dateAsString string

	dateAsString = strings.Trim(dateAsString, `"`)
	date, err := time.Parse("02-01-2006", dateAsString)
	if err != nil {
		return fmt.Errorf("Error parsing my-date: %s: %s", dateAsString, err)
	}
	*t = MyDate{
		UnderlyingTime: date,
	}
	return nil
}
