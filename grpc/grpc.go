package grpc

//
// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
//
// 	"github.com/charitan-go/email-server/internal/email/service"
// 	"github.com/charitan-go/email-server/pkg/proto"
// 	consulapi "github.com/hashicorp/consul/api"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/health"
// 	"google.golang.org/grpc/health/grpc_health_v1"
// )
//
// type GrpcServer struct {
// 	proto.UnimplementedEmailGrpcServiceServer
// 	grpcServer *grpc.Server
//
// 	emailSvc service.EmailService
// }
//
// func NewGrpcServer(emailSvc service.EmailService) *GrpcServer {
// 	grpcServer := grpc.NewServer()
// 	emailGrpcServer := &GrpcServer{}
//
// 	proto.RegisterEmailGrpcServiceServer(grpcServer, emailGrpcServer)
// 	emailGrpcServer.emailSvc = emailSvc
// 	emailGrpcServer.grpcServer = grpcServer
//
// 	address := os.Getenv("SERVICE_ID")
// 	grpcServiceName := fmt.Sprintf("%s-grpc", address)
//
// 	// Setup health server
// 	healthServer := health.NewServer()
// 	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
// 	healthServer.SetServingStatus(grpcServiceName, grpc_health_v1.HealthCheckResponse_SERVING)
//
// 	return emailGrpcServer
// }
//
// func (*GrpcServer) setupServiceRegistry() {
// 	log.Println("Start for grpc service registry")
//
// 	config := consulapi.DefaultConfig()
// 	config.Address = os.Getenv("SERVICE_REGISTRY_URI")
// 	consul, err := consulapi.NewClient(config)
// 	if err != nil {
// 		log.Println("Cannot connect with service registry", err)
// 	}
//
// 	address := os.Getenv("ADDRESS")
//
// 	grpcServiceId := fmt.Sprintf("%s-grpc", address)
// 	grpcRegistration := &consulapi.AgentServiceRegistration{
// 		ID:      grpcServiceId,
// 		Name:    grpcServiceId,
// 		Address: address,
// 		Port:    50051,
// 		Tags:    []string{"grpc"},
// 		Check: &consulapi.AgentServiceCheck{
// 			GRPC:     fmt.Sprintf("%v:%d", address, 50051),
// 			Interval: "10s",
// 			Timeout:  "5s",
// 			// DeregisterCriticalServiceAfter: "30s",
// 		},
// 	}
//
// 	err = consul.Agent().ServiceRegister(grpcRegistration)
// 	if err != nil {
// 		log.Fatalf("Failed to register gRPC service with Consul: %v", err)
// 	} else {
// 		log.Println("Register grpc service successfully")
// 	}
// }
//
// func (s *GrpcServer) Run() {
// 	s.setupServiceRegistry()
// 	log.Println("Setup service registry for grpc service ok")
//
// 	// Generate email pairs
// 	// s.emailSvc.GenerateEmailPairs()
//
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
//
// 	log.Println("GRPC server listening on :50051")
// 	if err := s.grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
//
// }

// func (s *GrpcServer) GetPrivateEmail(
// 	ctx context.Context,
// 	reqDto *proto.GetPrivateEmailRequestDto,
// ) (*proto.GetPrivateEmailResponseDto, error) {
// 	resDto, err := s.emailSvc.HandleGetPrivateEmailGrpc(reqDto)
// 	return resDto, err
// }
//
// func (s *GrpcServer) GetPublicEmail(
// 	ctx context.Context,
// 	reqDto *proto.GetPublicEmailRequestDto,
// ) (*proto.GetPublicEmailResponseDto, error) {
// 	resDto, err := s.emailSvc.HandleGetPublicEmailGrpc(reqDto)
// 	return resDto, err
// }
