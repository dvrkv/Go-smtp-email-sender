package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

func sendMail(subject string, templatePath string, to []string) {
	//get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct {
		Name  string
		Today time.Time
	}{Name: "dear subscriber!", Today: time.Now()})

	if err != nil {
		fmt.Println(err)
	}
	godotenv.Load(".env")
	auth := smtp.PlainAuth(
		"",
		os.Getenv("email"),
		os.Getenv("password"),
		"smtp.gmail.com",
	)
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()
	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		os.Getenv("email"), //letter sender
		to,                 //letter recipient
		[]byte(msg),        //whole message
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	recipients := []string{"testdvrkv@gmail.com"}
	sendMail(
		"Some newsletter",
		"./static/htmlBody.html",
		recipients,
	)
}
