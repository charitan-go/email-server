package rabbitmq

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitmqService interface {
	ConnectRabbitmq() (*amqp.Channel, error)
	DeclareExchange(ch *amqp.Channel, exchangeName string) error
	DeclareQueue(ch *amqp.Channel, queueName string) error
	QueueBind(ch *amqp.Channel, queueName, routingKey, exchangeName string) error

	Publish(ch *amqp.Channel, exChangeName, routingKey string, msg amqp.Publishing) error
	Consume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error)
}

type rabbitmqServiceImpl struct {
}

func NewRabbitmqService() RabbitmqService {
	return &rabbitmqServiceImpl{}
}

func (*rabbitmqServiceImpl) ConnectRabbitmq() (*amqp.Channel, error) {
	log.Println("In function startRabbitmqConsumer")

	amqpConnectionStr := fmt.Sprintf("amqp://%s:%s@message-broker:5672",
		os.Getenv("MESSAGE_BROKER_USER"),
		os.Getenv("MESSAGE_BROKER_PASSWORD"))
	conn, err := amqp.Dial(amqpConnectionStr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return nil, err
	}
	// defer ch.Close()

	return ch, nil
}

func (*rabbitmqServiceImpl) DeclareExchange(ch *amqp.Channel, exchangeName string) error {
	// exchangeName := "GET_PRIVATE_KEY"
	err := ch.ExchangeDeclare(
		exchangeName, // exchange name
		"topic",      // type: topic
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err != nil {
		log.Fatalf("Declare exchange failed")
	}

	return err

}

func (*rabbitmqServiceImpl) DeclareQueue(ch *amqp.Channel, queueName string) error {
	_, err := ch.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	return err
}

func (*rabbitmqServiceImpl) QueueBind(ch *amqp.Channel, queueName, routingKey, exchangeName string) error {
	err := ch.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}

	return err
}

func (*rabbitmqServiceImpl) Publish(ch *amqp.Channel, exchangeName, routingKey string, msg amqp.Publishing) error {
	err := ch.Publish(
		exchangeName, // exchange
		routingKey,   // routing key
		false,        // mandatory
		false,        // immediate
		msg,          // body
	)

	if err != nil {
		log.Fatalf("Cannot publish topic with exchangeName=%s, routingKey=%s\n", exchangeName, routingKey)
	}

	return err
}

func (*rabbitmqServiceImpl) Consume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		queueName, // queue name
		"",        // consumer tag
		true,      // auto-acknowledge
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return nil, err
	}

	return msgs, nil
}
