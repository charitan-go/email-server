package service

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"

	apigateway "github.com/charitan-go/key-server/external/api-gateway"
	"github.com/charitan-go/key-server/external/auth"
	"github.com/charitan-go/key-server/pkg/proto"
)

type KeyService interface {
	getPrivateKeyStr() string

	GenerateKeyPairs() error

	// GRPC Listener
	HandleGetPrivateKeyGrpc(*proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error)
	HandleGetPublicKeyGrpc(*proto.GetPublicKeyRequestDto) (*proto.GetPublicKeyResponseDto, error)
}

type keyServiceImpl struct {
	privateKey                 *rsa.PrivateKey
	publicKey                  *rsa.PublicKey
	authRabbitmqProducer       auth.AuthRabbitmqProducer
	apiGatewayRabbitmqProducer apigateway.ApiGatewayRabbitmqProducer
}

func NewKeyService(
	authRabbitmqProducer auth.AuthRabbitmqProducer,
	apiGatewayRabbitmqProducer apigateway.ApiGatewayRabbitmqProducer) KeyService {
	return &keyServiceImpl{nil, nil, authRabbitmqProducer, apiGatewayRabbitmqProducer}
}

func (svc *keyServiceImpl) getPrivateKeyStr() string {
	// Convert the RSA private key to DER-encoded PKCS#1 format
	privateKeyDER := x509.MarshalPKCS1PrivateKey(svc.privateKey)

	// Create a PEM block with the DER-encoded key
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyDER,
	}

	// Encode the PEM block to a []byte
	privateKeyPEM := pem.EncodeToMemory(pemBlock)

	// Convert the PEM []byte to a string
	privateKeyString := string(privateKeyPEM)
	return privateKeyString
}

func (svc *keyServiceImpl) getPublicKeyStr() string {
	// Marshal the public key to DER-encoded PKIX format.
	derBytes, _ := x509.MarshalPKIXPublicKey(svc.publicKey)

	// Create a PEM block with type "PUBLIC KEY".
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derBytes,
	}

	// Encode the PEM block to a memory buffer and return it as a string.
	return string(pem.EncodeToMemory(block))
}

func (svc *keyServiceImpl) GenerateKeyPairs() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Cannot generate private key: %v\n", err)
		return err
	}

	// Store both keys
	svc.privateKey = privateKey
	svc.publicKey = &privateKey.PublicKey

	// Send noti to auth server for get key pairs
	err = svc.authRabbitmqProducer.NotiGetPrivateKey()
	if err != nil {
		log.Fatalf("Cannot send noti to auth server: %v\n", err)
	} else {
		log.Println("Send noti to auth server success")
	}

	err = svc.apiGatewayRabbitmqProducer.NotiGetPublicKey()
	if err != nil {
		log.Fatalf("Cannot send noti to api gateway: %v\n", err)
	} else {
		log.Println("Send noti to api gateway success")
	}

	return nil
}

func (svc *keyServiceImpl) HandleGetPrivateKeyGrpc(*proto.GetPrivateKeyRequestDto) (*proto.GetPrivateKeyResponseDto, error) {
	if svc.privateKey == nil {
		return nil, fmt.Errorf("Private key not available")
	}

	return &proto.GetPrivateKeyResponseDto{PrivateKey: svc.getPrivateKeyStr()}, nil
}

func (svc *keyServiceImpl) HandleGetPublicKeyGrpc(*proto.GetPublicKeyRequestDto) (*proto.GetPublicKeyResponseDto, error) {
	if svc.publicKey == nil {
		return nil, fmt.Errorf("Public key not available")
	}

	return &proto.GetPublicKeyResponseDto{PublicKey: svc.getPublicKeyStr()}, nil
}
