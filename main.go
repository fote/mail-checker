package main

import (
	"fmt"
	"log"
	"net/smtp"
	"flag"
)

var username string
var password string

func buildMessage(from string, to string, subj string, body string) []byte {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", from)
	message += fmt.Sprintf("Subject: %s\r\n", subj)
	message += "\r\n" + body

	return []byte(message)
}

func main() {
	var host, port, from,to, subj, body string

	flag.StringVar(&host, "mailHost", "127.0.0.1", "Mail server address")
	flag.StringVar(&port, "mailPort", "25", "Mail server port")
	flag.StringVar(&password, "password", "password", "Password")
	flag.StringVar(&username, "username", "noreply", "From")
	flag.StringVar(&from, "from", "noreply@mydomain.ru", "From")
	flag.StringVar(&to, "to", "mymail@yandex.ru", "To")
	flag.StringVar(&subj, "subject", "Test message", "Subject")
	flag.StringVar(&body, "body", "This is a test. Ignore this message", "Body")

	flag.Parse()

	messageBody := buildMessage(from, to, subj, body)


	// Set up authentication information.
	auth := smtp.PlainAuth("", username, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, []string{to}, messageBody)
	if err != nil {
		log.Fatal(err)
	}

}
