package testit

// START OMIT
import (
	"fmt"
	"testing"
)

// function under test defined in other source file
// Naming convention:  starts with Test and has "t *testing.T" as parameter
func TestReverseAscii(t *testing.T) {
	value := Reverse("cram")
	if value != "marc" {
		t.Errorf("ERROR : Expecting[%s] Received[%s]", "marc", value)
	}
}

func ExampleReverse() {
	fmt.Println(Reverse("hello"))
	// Output: olleh
}

// END OMIT
