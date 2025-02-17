package app

import (
	"log"

	"github.com/charitan-go/email-server/grpc"
	"github.com/charitan-go/email-server/internal/email"
	"github.com/charitan-go/email-server/rabbitmq"
	"go.uber.org/fx"
)

// Run both servers concurrently
func runServers(grpcSrv *grpc.GrpcServer) {
	log.Println("In invoke")

	// Start gRPC server
	go func() {
		log.Println("In goroutine of grpc")
		grpcSrv.Run()
	}()
}

func Run() {

	fx.New(
		email.EmailModule,
		rabbitmq.RabbitmqModule,
		grpc.GrpcModule,
		fx.Invoke(runServers),
	).Run()
}
