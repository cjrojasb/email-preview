package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

type EmailTemplate2 struct {
	Welcome string
	Name    string
	Email   string
	Message string
}

func SendEmail() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GMAIL_USERNAME := os.Getenv("GMAIL_USERNAME")
	GMAIL_PASSWORD := os.Getenv("GMAIL_PASSWORD")
	gmailAuth := smtp.PlainAuth("", GMAIL_USERNAME, GMAIL_PASSWORD, "smtp.gmail.com")

	t, err := template.ParseFiles("./templates/email-template.html")
	if err != nil {
		log.Fatal(err)
	}

	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html;"
	body.Write([]byte(fmt.Sprintf("Subject: yourSubject\n%s\n\n", headers)))

	t.Execute(&body, EmailTemplate2{
		Welcome: "Dear cjrojasb",
		Name:    "Carlos",
		Email:   GMAIL_USERNAME,
		Message: "This is a test email",
	})
	if err != nil {
		log.Fatal(err)
	}

	smtp.SendMail("smtp.gmail.com:587", gmailAuth, GMAIL_USERNAME, []string{
		GMAIL_USERNAME}, body.Bytes())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully!")
}
