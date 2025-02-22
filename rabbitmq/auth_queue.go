package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	EMAIL_AUTH_QUEUE_NAME                = "email.auth.queue"
	EMAIL_AUTH_ROUTING_KEY               = "email.auth.*"
	DONOR_REGISTER_ACCOUNT_ROUTING_KEY   = "email.auth.donor.register_account"
	CHARITY_REGISTER_ACCOUNT_ROUTING_KEY = "email.auth.charity.register_account"
)

func (srv *RabbitmqServer) setupEmailAuthQueue(ch *amqp.Channel) (<-chan amqp.Delivery, error) {

	// Declare a queue for receiving topics from auth
	err := srv.rabbitmqSvc.DeclareQueue(ch, EMAIL_AUTH_QUEUE_NAME)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return nil, err
	}

	// Declare a queue for biding with topics from auth
	err = srv.rabbitmqSvc.QueueBind(ch, EMAIL_AUTH_QUEUE_NAME, EMAIL_AUTH_ROUTING_KEY, EMAIL_EXCHANGE_NAME)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return nil, err
	}

	// Consume messages from the queue.
	msgs, err := srv.rabbitmqSvc.Consume(ch, EMAIL_AUTH_QUEUE_NAME)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		return nil, err
	}

	return msgs, nil
}
