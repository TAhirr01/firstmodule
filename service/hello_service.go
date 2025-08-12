package service

import (
	"context"
	"fmt"

	"github.com/TAhirr01/firstmodule/pb"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func NewHelloService() pb.HelloServiceServer {
	return &HelloService{}
}

func (service *HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("SayHello called")
	helloResponse := "Salam " + req.GetName()

	return &pb.HelloResponse{
		HelloName: helloResponse,
	}, nil
}
