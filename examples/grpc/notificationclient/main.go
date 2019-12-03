
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
	reps = 10
)

func main() {
	client := New()
	err := client.Open(address)
	if err != nil {
		log.Fatalf("*** Error opening client: %v", err)
	}
	defer client.Close()


	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for idx:=0;idx<reps ; idx++ {
		uid, err := client.SendEmail(ctx, "mgrol@xebia.com",  fmt.Sprintf("My subject %d", idx), "My body")
		if err != nil {
			log.Printf("*** Error sending email: %v", err)
		} else {
			log.Printf("SendEmail-response: %+v", uid)
			status, err := client.GetStatus(ctx, uid)
			if err != nil {
				log.Printf("*** Error getting status of sms: %v", err)
			} else {
				log.Printf("GetStatus-response on email with uid %s: %+v", uid, status)
			}
		}
	}

	for idx:=0;idx<reps ; idx++ {
		uid, err := client.SendSms(ctx, "+31648928856", fmt.Sprintf("My body %d", idx))
		if err != nil {
			log.Printf("*** Error sending sms: %v", err)
		} else {
			log.Printf("SendSms-response: %+v", uid)

			status, err := client.GetStatus(ctx, uid)
			if err != nil {
				log.Printf("*** Error getting status of sms: %v", err)
			} else {
				log.Printf("GetStatus-response on sms with uid %s: %+v", uid, status)
			}
		}
	}
}