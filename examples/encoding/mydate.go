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
	formattedDate := t.UnderlyingTime.Format("02-01-2006")
	return json.Marshal(formattedDate)
}

func (t *MyDate) UnmarshalJSON(data []byte) error {
	var dateAsString string
	err := json.Unmarshal(data, &dateAsString)
	if err != nil {
		return fmt.Errorf("Error decoding json my-date as string: %+v: %s", data, err)
	}

	dateAsString = strings.Trim(dateAsString, `"`)
	date, err := time.Parse("02-01-2006", dateAsString)
	if err != nil {
		return fmt.Errorf("Error parsing json my-date: %s: %s", dateAsString, err)
	}
	*t = MyDate{
		UnderlyingTime: date,
	}
	return nil
}

func (t MyDate) MarshalText() ([]byte, error) {
	formattedDate := t.UnderlyingTime.Format("02-01-2006")
	return []byte(formattedDate), nil
}

func (t *MyDate) UnmarshalText(data []byte) error {
	var dateAsString string = string(data)
	date, err := time.Parse("02-01-2006", dateAsString)
	if err != nil {
		return fmt.Errorf("Error parsing xml-my-date: %s: %s", dateAsString, err)
	}
	*t = MyDate{
		UnderlyingTime: date,
	}
	return nil
}
