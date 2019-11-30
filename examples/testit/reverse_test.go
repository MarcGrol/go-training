package testit

// START OMIT
import (
	"testing"
)

// Naming convention:  starts with Test and has "t *testing.T" as parameter
func TestReverseAscii(t *testing.T) {
	value := Reverse("cram")
	if value != "marc" {
		t.Errorf("ERROR : Expecting[%s] Received[%s]", "marc", value)
	}
}

// END OMIT

//func ExampleReverse() {
//	fmt.Println(Reverse("hello"))
//	// Output: olleh
//}
