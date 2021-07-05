package myuuid

import "time"

//START OMIT
func New() string {
	now := time.Now()
	return now.Format(time.RFC3339)
}
//END OMIT