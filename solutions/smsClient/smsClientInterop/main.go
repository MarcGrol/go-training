package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/MarcGrol/go-training/solutions/smsClient"
	"os"
)

var PROGRAM = os.Args[0]

func main() {
	cfg := parseArgs()

	c := context.Background()

	err := smsClient.NewTwilloSmsClient(cfg.accountId, cfg.password).SendSms(c, cfg.recipientPhoneNumber, cfg.messageText)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending sms via twillo: %s\n", err)
		os.Exit(-1)
	}
	os.Exit(0)
}

type config struct{
	accountId string
	password string
	recipientPhoneNumber string
	messageText string
}

func parseArgs() config {

	help := flag.Bool("help", false, "This help text")
	to := flag.String("to", "+31648928856", "Recipient phone number")
	msg := flag.String("msg", "Test sms", "Sms contents")
	account := flag.String("account", "", "Twilio accountId")
	password := flag.String("password", "", "Twilio password")

	flag.Parse()

	if help != nil && *help {
		fmt.Fprintf(os.Stderr, "\nUsage:\n")
		fmt.Fprintf(os.Stderr, " %s [flags]\n", PROGRAM)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
		os.Exit(-1)
	}

	return config{
		accountId:*account,
		password:*password,
		recipientPhoneNumber:*to,
		messageText:*msg,
	}
}
