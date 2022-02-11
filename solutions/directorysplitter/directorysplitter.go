package directoryplitter

import "strings"

func DirectorySplitter(randomString string) []string {
	return strings.Split(randomString, "/")
}
