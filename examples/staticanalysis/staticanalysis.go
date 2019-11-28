package staticanalysis

import (
	"io/ioutil"
	"log"
)

type person struct { // unused
  Name string
  Age int
}

func WriteMessage(msg string) { // Exported function not documented
	ioutil.WriteFile("myfile.txt", []byte(msg), 0) // No error handling
	log.Printf("Did it: %s") // Missing vararg param
}
