package main

import (
	"fmt"
	"time"
)

// START OMIT
func sendMsg(c chan string) { // HL
	time.Sleep(50 * time.Millisecond) // HL
	c <- "hi there"                   // HL
} // HL

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(300 * time.Millisecond)
	msgChannel := make(chan string) // HL
	go sendMsg(msgChannel)          // HL
	for {
		select {
		case msg := <-msgChannel:
			fmt.Printf("msg: %s\n", msg)
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		}
	}
}

// END OMIT
