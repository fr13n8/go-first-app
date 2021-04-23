package models

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func (u User) SendVerifyEmail(c *gin.Context) error {
	var (
		from     string = os.Getenv("EMAIL_NAME")
		password string = os.Getenv("EMAIL_PASSWORD")
		smtpHost string = os.Getenv("EMAIL_HOST")
		smtpPort string = os.Getenv("EMAIL_PORT")
	)

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		return err
	}

	to := []string{
		u.Email,
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error getting working directory")
		return err
	}

	token, err := u.JwtGenerate(c)
	if err != nil {
		log.Println("Error generating token")
		return err
	}

	AppHost := fmt.Sprintf("%s%s", os.Getenv("APP_PROTOCOL"), os.Getenv("APP_HOST"))
	AppPort := os.Getenv("APP_PORT")
	URL := fmt.Sprintf("%s:%s", AppHost, AppPort)

	verifyLink := fmt.Sprintf("%s/auth/verify/%s", URL, token)

	t, err := template.ParseFiles(filepath.Join(dir, "templates", "email", "verify-email.html"))
	if err != nil {
		log.Println("Error parsing email template")
		return err
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a xuy subject \n%s\n\n", mimeHeaders)))

	message := "Please verify your account"

	if err := t.Execute(&body, struct {
		Name      string
		Message   string
		VerifyURL string
	}{
		Name:      u.Name,
		Message:   message,
		VerifyURL: verifyLink,
	}); err != nil {
		log.Println("Error executing message template")
		return err
	}

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send Email
	if err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes()); err != nil {
		log.Println("Error sending email")
		return err
	}

	log.Printf("Verify mail to %s sended.", u.Email)

	return nil
}
