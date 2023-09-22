package main

import (
	"context"
	"fmt"
	"net"
	"sample/config"
	"sample/constants"
	"sample/controllers"
	pb "sample/proto"
	"sample/services"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initApp(client *mongo.Client) {
	ctx := context.TODO()
	mcoll := client.Database(constants.DatabaseName).Collection("users")
	service := services.NewUserService(ctx, mcoll)
	controllers.UserService = service
}

func main() {
	fmt.Println("Starting server...")
	client, err := config.ConnectDataBase()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.TODO())
	initApp(client)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &controllers.RPCServer{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
