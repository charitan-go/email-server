package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	"github.com/charitan-go/email-server/pkg/proto"
)

type EmailService interface {
	getPrivateEmailStr() string

	GenerateEmailPairs() error

	// GRPC Listener
	HandleGetPrivateEmailGrpc(*proto.GetPrivateEmailRequestDto) (*proto.GetPrivateEmailResponseDto, error)
	HandleGetPublicEmailGrpc(*proto.GetPublicEmailRequestDto) (*proto.GetPublicEmailResponseDto, error)
}

type emailServiceImpl struct {
	privateEmail                 *rsa.PrivateEmail
	publicEmail                  *rsa.PublicEmail
	authRabbitmqProducer       auth.AuthRabbitmqProducer
	apiGatewayRabbitmqProducer apigateway.ApiGatewayRabbitmqProducer
}

func NewEmailService(
	authRabbitmqProducer auth.AuthRabbitmqProducer,
	apiGatewayRabbitmqProducer apigateway.ApiGatewayRabbitmqProducer) EmailService {
	return &emailServiceImpl{nil, nil, authRabbitmqProducer, apiGatewayRabbitmqProducer}
}

func (svc *emailServiceImpl) getPrivateEmailStr() string {
	// Convert the RSA private email to DER-encoded PKCS#1 format
	privateEmailDER := x509.MarshalPKCS1PrivateEmail(svc.privateEmail)

	// Create a PEM block with the DER-encoded email
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateEmailDER,
	}

	// Encode the PEM block to a []byte
	privateEmailPEM := pem.EncodeToMemory(pemBlock)

	// Convert the PEM []byte to a string
	privateEmailString := string(privateEmailPEM)
	return privateEmailString
}

func (svc *emailServiceImpl) getPublicEmailStr() string {
	// Marshal the public email to DER-encoded PKIX format.
	derBytes, _ := x509.MarshalPKIXPublicEmail(svc.publicEmail)

	// Create a PEM block with type "PUBLIC KEY".
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	}

	// Encode the PEM block to a memory buffer and return it as a string.
	return string(pem.EncodeToMemory(block))
}

func (svc *emailServiceImpl) GenerateEmailPairs() error {
	privateEmail, err := rsa.GenerateEmail(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Cannot generate private email: %v\n", err)
		return err
	}

	// Store both emails
	svc.privateEmail = privateEmail
	svc.publicEmail = &privateEmail.PublicEmail

	// Send noti to auth server for get email pairs
	err = svc.authRabbitmqProducer.NotiGetPrivateEmail()
	if err != nil {
		log.Fatalf("Cannot send noti to auth server: %v\n", err)
	} else {
		log.Println("Send noti to auth server success")
	}

	err = svc.apiGatewayRabbitmqProducer.NotiGetPublicEmail()
	if err != nil {
		log.Fatalf("Cannot send noti to api gateway: %v\n", err)
	} else {
		log.Println("Send noti to api gateway success")
	}

	return nil
}

func (svc *emailServiceImpl) HandleGetPrivateEmailGrpc(*proto.GetPrivateEmailRequestDto) (*proto.GetPrivateEmailResponseDto, error) {
	if svc.privateEmail == nil {
		return nil, fmt.Errorf("Private email not available")
	}

	return &proto.GetPrivateEmailResponseDto{PrivateEmail: svc.getPrivateEmailStr()}, nil
}

func (svc *emailServiceImpl) HandleGetPublicEmailGrpc(*proto.GetPublicEmailRequestDto) (*proto.GetPublicEmailResponseDto, error) {
	if svc.publicEmail == nil {
		return nil, fmt.Errorf("Public email not available")
	}

	return &proto.GetPublicEmailResponseDto{PublicEmail: svc.getPublicEmailStr()}, nil
}
