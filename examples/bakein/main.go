package main

import "fmt"

var (
	VersionHash string
)

// to run this:
// go install -ldflags="-X 'main.VersionHash=$(git rev-parse HEAD)'" && bakein or
// go build -ldflags="-X 'main.VersionHash=$(git rev-parse HEAD)'" && ./bakein
//
// replace "main" with import name when applied to libraries
func main() {
	fmt.Printf("VersionHash:%s\n", VersionHash)
}
