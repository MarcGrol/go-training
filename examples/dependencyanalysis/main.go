// This package is used to demonstrate use of documentaion
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "Hello %s", "there")
}
