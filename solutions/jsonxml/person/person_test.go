package person

import (
	// include standard testing-package
	"bytes"
	"reflect"
	"testing"
)

func TestJson(t *testing.T) {
	before := Person{
		Name:         "Marc",
		Age:          42,
		Interests:    []string{"Running", "Cycling", "Hockey"},
		privateField: false}

	// this log only appears in verbose test: go test -v
	t.Logf("before:%+v\n", before)

	// Convert to json
	var buffer bytes.Buffer
	err := before.ToJson(&buffer)
	if err != nil {
		t.Errorf("Expected json encoding to succeed: %+v", err)
	}

	t.Logf("json:%s", buffer.Bytes())

	// Read back json into struct
	after, err := FromJson(&buffer)
	if err != nil {
		t.Errorf("Expected json decoding to succeed: %+v", err)
	}

	t.Logf("after:%+v\n", after)

	// compare 'before' and 'after' structs
	equal := reflect.DeepEqual(before, *after)
	if equal != true {
		t.Errorf("Expected '%+v' not equal to actual: '%+v", before, after)
	}
}
