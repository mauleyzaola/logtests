package main

import (
	"github.com/mauleyzaola/logtests/pkg/senders"
)

func main() {}

func Send(sender senders.EmailSender, email string) error {
	return sender.Send(email)
}
