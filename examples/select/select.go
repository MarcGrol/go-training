package main

import (
	"fmt"
	"time"
)

// START OMIT
func sendMsg(c chan<- string) { // chan<-: Only allows writing to
	time.Sleep(100 * time.Millisecond)
	c <- "Put your helmet on"
}

func main() {
	tickerChannel := time.Tick(800 * time.Millisecond) // Emit a signal every ...
	afterChannel := time.After(3 * time.Second)        // Fires once after ...
	msgChannel := make(chan string)                    // Used by sendMsg to return its result
	go sendMsg(msgChannel)
	for {
		select { // blocking untill msg received on one of its channels
		case msg := <-msgChannel:
			fmt.Printf("msg: %s\n", msg) // stay in loop
		case <-tickerChannel:
			fmt.Println("tickerChannel.") // stay in loop
		case <-afterChannel:
			fmt.Println("BOOM!")
			return // abort loop
		}
	}
}

// END OMIT
