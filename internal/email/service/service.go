package service

import (
	"log"

	"github.com/charitan-go/email-server/external/auth"
	"github.com/charitan-go/email-server/external/inbucket"
)

type EmailService interface {
	HandleRegisterDonorAccountRabbitmq(reqDto auth.SendRegisterDonorAccountEmailRequestDto) error
}

type emailServiceImpl struct {
	inbucketService inbucket.InbucketService
}

func NewEmailService(inbucketService inbucket.InbucketService) EmailService {
	return &emailServiceImpl{inbucketService}
}

// HandleRegisterAccountRabbitmq implements EmailService.
func (svc *emailServiceImpl) HandleRegisterDonorAccountRabbitmq(reqDto auth.SendRegisterDonorAccountEmailRequestDto) error {
	log.Println("Receive topic for register account")

	// Send email
	// to := []string{reqDto.Email}

	// Load body from template engine

	// Build the email message (including headers)
	// subject := "Subject: Test Email from Inbucket\n"
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

	return nil
}
