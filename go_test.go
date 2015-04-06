package main

// START OMIT
import (
	"testing"
)

// function under test defined in other source file
func TestReverseAscii(t *testing.T) {
	value := reverse("cram")
	if value != "marc" {
		t.Errorf("ERROR : Expecting[%s] Received[%s]", "marc", value)
	}
}

// END OMIT

func reverse(in string) string {
	return in
}

func main() {
}
