package tests

import (
	"fmt"
	"testing"
)

// START OMIT
func TestValidateEmail(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "", want: false},
		{input: "@", want: false},
		{input: "@xebia", want: false},
		{input: "m/grol@xebia..com", want: false},
		{input: "mgrol@xebia..com", want: false},
		{input: "mgrol@xebia", want: true},
		{input: "mgrol+dev@xebia", want: true},
		{input: "marc.grol@gmail.com", want: true},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Testcase: %s", tc.input), func(t *testing.T) {
			got := IsValidEmailAddress(tc.input)
			if got != tc.want {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}

// END OMIT
