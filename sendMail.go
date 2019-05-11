package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func sendMail(emailCred Email) {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		emailCred.From,
		emailCred.Pwd,
		emailCred.SMTP,
	)

	msg := "From: " + emailCred.From + "\n" +
		"To: " + emailCred.To + "\n" +
		"Subject: " + emailCred.Subject + "\n\n" +
		"I'm home!"

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	server := fmt.Sprintf("%s:%d", emailCred.SMTP, emailCred.Port)
	err := smtp.SendMail(
		server,
		auth,
		emailCred.From,
		[]string{emailCred.To},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal(err)
	}
}
