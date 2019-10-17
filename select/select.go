package main

import (
	"fmt"
	"time"
)

// START OMIT
func sendMsg(c chan string) { // HL
	time.Sleep(100 * time.Millisecond) // HL
	c <- "Put your helmet on"          // HL
} // HL

func main() {
	tick := time.Tick(800 * time.Millisecond)
	boom := time.After(3 * time.Second)
	msgChannel := make(chan string)
	go sendMsg(msgChannel)
	for {
		select { // blocking untill msg received on one of its channels
		case msg := <-msgChannel:
			fmt.Printf("msg: %s\n", msg) // stay in loop
		case <-tick:
			fmt.Println("tick.") // stay in loop
		case <-boom:
			fmt.Println("BOOM!")
			return // abort loop
		}
	}
}

// END OMIT
