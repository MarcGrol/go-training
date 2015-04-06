package main

import (
	"fmt"
)

const myConstString = "golang"

func main() {
	var status bool // uninitialized
	var idx int64 = 42
	i := 256 // := derives type from right-side
	longString := `{
		"why":"Usefull to embed json in source"
	}`
	myList := []int{1, 3, 5, 7}
	fmt.Printf(
		"status:%v\nidx:%d\ni=%d\nmy-const-string:%s\nmy-long-string:%s\nmy-list:%+v\n",
		status, idx, i, myConstString, longString, myList)
}
