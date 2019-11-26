package main

import (
	"fmt"
)

// START OMIT
const myConstString = "golang"

func main() {
	var status bool // uninitialized -> default (=false)
	var idx int64 = 42
	// := short notation: derives type from right-hand-side // HL
	i := 256 // HL
	longString := `{
		"why":"Usefull to embed json in source"
	}`
	myList := []int{1, 3, 5, 7}
	fmt.Printf(
		"status:%v\nidx:%d\ni=%d\nmy-const-string:%s\nmy-long-string:%s\nmy-list:%+v\n",
		status, idx, i, myConstString, longString, myList)
}

// END OMIT
