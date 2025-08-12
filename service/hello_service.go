package service

import (
	"context"
	"firstmodule/pb"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func NewHelloService() pb.HelloServiceServer {
	return &HelloService{}
}

func (service *HelloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	helloResponse := "Salam " + req.GetName()

	return &pb.HelloResponse{
		HelloName: helloResponse,
	}, nil
}
