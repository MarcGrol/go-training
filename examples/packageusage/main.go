	package main

import (
	"fmt"  // package from stdlib
	"os"   // package from stdlib
	"time" // package from stdlib

	"github.com/google/uuid" // third-party package
)

func main() {
	u := uuid.New()   // use package-name as prefix
	now := time.Now() // use package-name as prefix

	fmt.Fprintf(os.Stdout, "uuid: %s\ntime: %s", u.String(),
		now.Format(time.RFC3339))
}
