package main

import (
	"context"
	"log"

	"github.com/TAhirr01/firstmodule/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("cannot connect to grpc server")
	}
	client := pb.NewUserServiceClient(lis)

	//registerResponse, err := client.RegisterUser(context.Background(), &pb.UserRequest{
	//	Name:     "Test",
	//	Email:    "tahirqasimov002@gmail.com",
	//	Password: "123456",
	//	Age:      222,
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(registerResponse)
	//log.Println()

	findbyidResponse, err := client.FindUserById(context.Background(), &pb.UserId{Id: 1})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(findbyidResponse)
	log.Println()

	finAllUsersResponse, err := client.FindAllUsers(context.Background(), &empty.Empty{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(finAllUsersResponse)
	log.Println()

}
