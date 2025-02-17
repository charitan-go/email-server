package app

import (
	"log"

	"github.com/charitan-go/email-server/internal/email"
	"github.com/charitan-go/email-server/rabbitmq"
	"go.uber.org/fx"
)

// Run both servers concurrently
func runServers(rabbitmqSrv *rabbitmq.RabbitmqServer) {
	log.Println("In invoke")

	// Start gRPC server
	go func() {
		log.Println("In goroutine of rabbitmq server")
		rabbitmqSrv.Run()

		// select {}
	}()
}

func Run() {

	fx.New(
		email.EmailModule,
		rabbitmq.RabbitmqModule,
		fx.Invoke(runServers),
	).Run()
}
