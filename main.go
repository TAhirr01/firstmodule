package main

import (
	"log"
	"net"

	initializers "github.com/TAhirr01/firstmodule/initializer"
	"github.com/TAhirr01/firstmodule/pb"
	"github.com/TAhirr01/firstmodule/service"

	"google.golang.org/grpc"
)

func init() {
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &service.UserService{})

	log.Println("grpc server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
