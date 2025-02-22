package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	// "time"

	"github.com/charitan-go/email-server/external/auth"
	email "github.com/charitan-go/email-server/internal/email/service"
	amqp "github.com/rabbitmq/amqp091-go"
)

var EMAIL_EXCHANGE_NAME string = "email.exchange"

type RabbitmqServer struct {
	rabbitmqSvc RabbitmqService
	emailSvc    email.EmailService
}

func NewRabbitmqServer(rabbitmqSvc RabbitmqService, emailSvc email.EmailService) *RabbitmqServer {
	return &RabbitmqServer{rabbitmqSvc, emailSvc}
}

func (srv *RabbitmqServer) startRabbitmqConsumer() error {
	// ch, err := srv.rabbitmqSvc.ConnectRabbitmq()
	log.Println("In function startRabbitmqConsumer")

	// time.Sleep(5 * time.Second)

	amqpConnectionStr := fmt.Sprintf("amqp://%s:%s@message-broker:5672",
		os.Getenv("MESSAGE_BROKER_USER"),
		os.Getenv("MESSAGE_BROKER_PASSWORD"))
	log.Printf("Connection str: %s\n", amqpConnectionStr)
	conn, err := amqp.Dial(amqpConnectionStr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare exchange name for private key
	err = srv.rabbitmqSvc.DeclareExchange(ch, EMAIL_EXCHANGE_NAME)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
		return err
	}

	// Setup queue from auth-server
	authMsgs, err := srv.setupEmailAuthQueue(ch)
	if err != nil {
		log.Fatalf("Setup auth queue failed %v\n", err)
		return err
	}

	forever := make(chan bool)
	go func() {
		log.Println("Inside the loop to process exchange topics")
		for d := range authMsgs {
			log.Printf("Receive topic %s\n", d.RoutingKey)
			switch d.RoutingKey {
			case REGISTER_DONOR_ACCOUNT_ROUTING_KEY:
				{
					log.Printf("Received message from exchange %s: %s\n", d.RoutingKey, d.Body)
					var reqDto auth.SendRegisterDonorAccountEmailRequestDto
					err := json.Unmarshal(d.Body, &reqDto)
					if err != nil {
						log.Fatalf("Cannot parse rabbitmq reqDto: %v\n", err)
						return
					}

					go srv.emailSvc.HandleRegisterDonorAccountRabbitmq(&reqDto)
				}
			case REGISTER_CHARITY_ACCOUNT_ROUTING_KEY:
				{
					// log.Printf("Received message from exchange %s: %s\n", d.RoutingKey, d.Body)
					// var reqDto auth.SendRegisterDonorAccountEmailRequestDto
					// err := json.Unmarshal(d.Body, &reqDto)
					// if err != nil {
					// 	log.Fatalf("Cannot parse rabbitmq reqDto: %v\n", err)
					// 	return
					// }
					//
					// srv.emailSvc.HandleRegisterDonorAccountRabbitmq(reqDto)

				}
			}
		}
	}()

	<-forever

	return nil

}

func (s *RabbitmqServer) Run() {
	// Start rabbitmq consumer
	s.startRabbitmqConsumer()
}
