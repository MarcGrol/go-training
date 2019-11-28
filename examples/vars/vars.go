package main

import (
	"fmt"
)

// START OMIT
const myConstString = "golang"

func main() {
	fmt.Printf("my-const-string: %s\n", myConstString)

	var status bool // uninitialized -> default (=false)
	fmt.Printf("status: %v\n", status)

	// := short notation: derives type from right-hand-side // HL
	idx := 256 // HL
	fmt.Printf("idx: %d\n", idx)

	longString := `{
		"why":"Usefull to embed json in source"
	}`
	fmt.Printf("my-long-string: %s\n", longString)
}

// END OMIT
