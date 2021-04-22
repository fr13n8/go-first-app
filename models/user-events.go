package models

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func (u *User) SendVerifyEmail() error {
	var (
		from     string = os.Getenv("EMAIL_NAME")
		password string = os.Getenv("EMAIL_PASSWORD")
		smtpHost string = os.Getenv("EMAIL_HOST")
		smtpPort string = os.Getenv("EMAIL_PORT")
	)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	to := []string{
		u.Email,
	}

	message := []byte("Please verify your account")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send Email
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message); err != nil {
		return err
	}

	fmt.Printf("Verify mail to %s sended", u.Email)
	return nil
}
