package main

import (
	"fmt"
	"strconv"
	"sync"
)

// START OMIT

type Currency struct {
	sync.Mutex // Embed notation // HL
	amount     float64
	code       string
}

func (c *Currency) Add(i float64) {
	c.Lock()         // HL
	defer c.Unlock() // HL
	c.amount += i
}

func (c *Currency) Display() string {
	c.Lock()         // HL
	defer c.Unlock() // HL
	return strconv.FormatFloat(c.amount, 'f', 2, 64) + " " + c.code
}

// END OMIT

func main() {
	c := Currency{amount: 10, code: "EU"}
	fmt.Printf("initial amount:%s\n", c.Display())
	c.Add(-7.54)
	fmt.Printf("amount after payment:%s\n", c.Display())
}
