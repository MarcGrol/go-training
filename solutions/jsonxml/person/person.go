package person

import (
	"encoding/json"
	"io"
)

type Person struct {
	// Fields must accessible by "encoding/json"-package
	Name         string   `json:"name"`
	Age          int      `json:"age"`
	Interests    []string `json:"interests"`
	NotInJson    bool     `json:"-"`
	privateField bool     // private field does not end up in json
}

func FromJson(reader io.Reader) (*Person, error) {
	person := Person{}
	err := json.NewDecoder(reader).Decode(&person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (p Person) ToJson(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(p)
}
