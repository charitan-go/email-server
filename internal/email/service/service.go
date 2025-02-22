package service

import (
	"bytes"
	"html/template"
	"log"

	"github.com/charitan-go/email-server/external/auth"
	"github.com/charitan-go/email-server/external/inbucket"
)

var EMAIL_TEMPLATE_DIR = "resource/email"

type EmailService interface {
	HandleRegisterDonorAccountRabbitmq(reqDto *auth.SendRegisterDonorAccountEmailRequestDto) error
}

type emailServiceImpl struct {
	inbucketService inbucket.InbucketService
}

func NewEmailService(inbucketService inbucket.InbucketService) EmailService {
	return &emailServiceImpl{inbucketService}
}

// HandleRegisterAccountRabbitmq implements EmailService.
func (svc *emailServiceImpl) HandleRegisterDonorAccountRabbitmq(reqDto *auth.SendRegisterDonorAccountEmailRequestDto) error {
	log.Println("Receive topic for register account")

	// Prepare toEmail
	toEmail := []string{reqDto.Email}

	// Load body from template engine
	tmpl, err := template.ParseFiles(EMAIL_TEMPLATE_DIR + "/register_donor_account.html")
	if err != nil {
		log.Fatalf("Cannot load email template: %v\n", err)
		return err
	}

	// Prepare email body
	var body bytes.Buffer

	// Build the email message (including headers)
	body.WriteString("Subject: Test Email from Inbucket\n")
	body.WriteString("MIME-version: 1.0;\r\n")
	body.WriteString("Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n")

	// Execute the template with the reqDto
	if err := tmpl.Execute(&body, reqDto); err != nil {
		log.Fatalf("Cannot execute the mail template: %v\n", err)
		return err
	}

	// Send email
	if err := svc.inbucketService.SendEmail(&inbucket.SendEmailRequestDto{ToEmail: toEmail, Content: body.Bytes()}); err != nil {
		log.Fatalf("Cannot send email: %v\n", err)
		return err
	}

	// mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	// body := "Hello,\n\nThis is a test email sent to Inbucket.\n\nRegards,\nYour App"
	// message := []byte(subject + mime + body)
	//
	// // Inbucket SMTP server address (default is 127.0.0.1:2500)
	// addr := "mail-server:2500"
	//
	// // Send the email without authentication (Inbucket accepts unauthenticated mail for testing)
	// err := smtp.SendMail(addr, nil, from, to, message)
	// if err != nil {
	// 	log.Fatalf("Failed to send email: %v", err)
	// }
	// log.Println("Email sent successfully!")
	//
	return nil
}
