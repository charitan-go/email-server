package apigateway

import (
	"fmt"
	"log"
	"os"

	"github.com/charitan-go/key-server/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ApiGatewayRabbitmqProducer interface {
	NotiGetPublicKey() error
}

type apiGatewayRabbitmqProducerImpl struct {
	rabbitmqSvc rabbitmq.RabbitmqService
}

func NewApiGatewayRabbitmqProducer(rabbitmqSvc rabbitmq.RabbitmqService) ApiGatewayRabbitmqProducer {
	return &apiGatewayRabbitmqProducerImpl{rabbitmqSvc}
}

func (p *apiGatewayRabbitmqProducerImpl) NotiGetPublicKey() error {
	amqpConnectionStr := fmt.Sprintf("amqp://%s:%s@message-broker:5672",
		os.Getenv("MESSAGE_BROKER_USER"),
		os.Getenv("MESSAGE_BROKER_PASSWORD"))
	conn, err := amqp.Dial(amqpConnectionStr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}
	defer ch.Close()

	// Declare a topic exchange (will create it if it doesn't exist)
	exchangeName := "key.exchange"
	err = p.rabbitmqSvc.DeclareExchange(ch, exchangeName)
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
		return err
	}

	// Publish a message indicating the key generation is complete
	routingKey := "key.get_public_key"
	body := "Public key generated successfully"
	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	}
	err = p.rabbitmqSvc.Publish(ch, exchangeName, routingKey, msg)
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	} else {
		log.Printf("Published message: %s", body)
	}

	return nil
}
