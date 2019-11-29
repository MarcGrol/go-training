package main

import "fmt"

// START OMIT
func main() {
	letters := []string{"a", "b", "c", "d"}
	fmt.Printf( "before: %v:   length: %d, capacity: %d (%p)\n",
		letters, len(letters), cap(letters), letters)

	// add items
	// append( letters, "e" ) // wrong!!! // HL
	letters = append( letters, "e" ) // why? realloc when npo longer fits // HL
	fmt.Printf( "after:  %v: length: %d, capacity: %d (%p)\n\n",
		letters, len(letters), cap(letters), letters) // pointer has changed

	// access items
	fmt.Printf("first:   %v\n", letters[0])              // a
	fmt.Printf("nothing: %v\n", letters[2:2])            // []
	fmt.Printf("begin:   %v\n", letters[:2])             // [a b]
	fmt.Printf("middlet: %v\n", letters[1:3])            // [b c]
	fmt.Printf("end:     %v\n", letters[3:])             // [d e]
	fmt.Printf("last:    %v\n", letters[len(letters)-1]) // e

	// iterate
	for idx, value := range letters {
		fmt.Printf("values[%d] = %s\n", idx, value)
	}
}
// END OMIT