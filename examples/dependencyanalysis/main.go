// This package is used to demonstrate use of documentaion
package main

import (
	"fmt"
	"os"
)

// go list -f '{{ .Name }}: {{ .Doc }}' fmt
// go list -f '{{ .Name }}: {{ .Doc }}' # current package
// go list -f '{{ join .Imports  "\n" }}' # Nice input for script

// go doc fmt
// go doc fmt Fprintf
//


func main() {
	fmt.Fprintf(os.Stdout,"Hello %s", "there")
}
