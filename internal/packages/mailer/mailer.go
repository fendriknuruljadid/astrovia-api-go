package mailer

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	from := os.Getenv("SMTP_EMAIL")
	pass := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth(
		"",
		from,
		pass,
		"smtp.gmail.com",
	)

	msg := []byte(fmt.Sprintf(
		"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"+
			"%s",
		subject,
		body,
	))

	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		[]string{to},
		msg,
	)
}

func SendInvoiceEmailAsync(to string, data map[string]any) {
	go func() {
		html, err := RenderTemplate("invoice.html", data)
		if err != nil {
			fmt.Println("render error:", err)
			return
		}

		if err := SendEmail(to, "Invoice Pembayaran", html); err != nil {
			fmt.Println("send email error:", err)
		}
	}()
}

func SendOTPEmail(to string, otp string) error {

	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	subject := "Your Verification Code"
	body := fmt.Sprintf("Your verification code is: %s\n\nThis code expires in 10 minutes.", otp)

	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	err := smtp.SendMail(
		smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{to},
		message,
	)

	return err
}
