package main

import (
	"context"
	"log"

	"github.com/TAhirr01/firstmodule/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("cannot connect to grpc server")
	}
	client := pb.NewUserServiceClient(lis)

	response, err := client.RegisterUser(context.Background(), &pb.UserRequest{
		Name:     "Tahir",
		Email:    "tahirqasimov001@gmail.com",
		Password: "123456",
		Age:      20,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(response)

}
