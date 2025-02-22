package inbucket

import (
	"log"
	"net/smtp"
)

const (
	SENDER_EMAIL = "no-reply@charitan-go.com"
	ADDRESS      = "mail-server:2500"
)

type InbucketService interface {
	SendEmail(*SendEmailRequestDto) error
}

type inbucketServiceImpl struct {
}

func NewInbucketService() InbucketService {
	return &inbucketServiceImpl{}
}

func (s *inbucketServiceImpl) SendEmail(reqDto *SendEmailRequestDto) error {

	if err := smtp.SendMail(ADDRESS, nil, SENDER_EMAIL, reqDto.ToEmail, reqDto.Content); err != nil {
		log.Fatalf("Cannot send email via smtp: %v\n", err)
		return err
	}

	return nil
}
