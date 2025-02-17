rm -rf pkg/proto/*.pb.go

protoc --go_out=. --go-grpc_out=. pkg/proto/key.proto
