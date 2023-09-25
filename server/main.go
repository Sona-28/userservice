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
	usercoll := client.Database(constants.DatabaseName).Collection("users")
	rolecoll := client.Database(constants.DatabaseName).Collection("roles")
	service := services.NewUserService(ctx, usercoll, rolecoll)
	controllers.UserService = service
}

func main() {
	fmt.Println("Starting server...")
	fmt.Println("hello")
	client, err := config.ConnectDataBase()
	if err != nil {
		fmt.Println("Error connecting to database")
		panic(err)
	}
	fmt.Println("Connected to database")

	defer client.Disconnect(context.TODO())
	initApp(client)
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &controllers.RPCServer{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
