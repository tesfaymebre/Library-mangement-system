package utils

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

// GenerateOTP generates a 6-digit OTP
func GenerateOTP(email string) (string, error) {
	otp := make([]byte, 6)
	_, err := rand.Read(otp)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(otp); i++ {
		otp[i] = (otp[i] % 10) + '0'
	}

	return string(otp), nil
}

// SendOTP sends the OTP to the specified email address
func SendOTP(email, otp string) error {
	from := os.Getenv("EMAIL_VERIFICATION_SENDER_EMAIL")
	if from == "" {
		return fmt.Errorf("EMAIL_VERIFICATION_SENDER_EMAIL not set in environment variables")
	}

	password := os.Getenv("EMAIL_VERIFICATION_SENDER_PASSWORD")
	if password == "" {
		return fmt.Errorf("EMAIL_VERIFICATION_SENDER_PASSWORD not set in environment variables")
	}

	to := email
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// create email subject
	emailSubject := "Verify Your Email"

	// generate email content
	emailContent := `
	<p>Thank you for signing up. To verify your account and complete the signup process, please use the following verification code:</p>
	<h3>` + otp + `</h3>
	<p><strong>This verification code is valid for 5 minutes.</strong> Please enter it on the verification page to proceed.</p>
	<p>If you did not sign up for an Akil account, please ignore this email.</p>`

	// create the email body
	emailBody := "<h1>" + emailSubject + "</h1><p>" + emailContent + "</p>"

	// create the email message
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + emailSubject + "\n" +
		"Content-Type: text/html; charset=UTF-8" + "\n\n" +
		emailBody

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
	log.Println("Email sent")
	if err != nil {
		log.Println("SMTP error:", err) // Add this to see the actual error
		return err
	}

	return nil
}
