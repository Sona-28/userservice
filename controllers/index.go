package controllers

import (
	"context"
	"fmt"
	"sample/interfaces"
	"sample/models"
	pb "sample/proto"
)

type RPCServer struct{
	pb.UnimplementedUserServer
}

var (
	UserService interfaces.IUser
)

func (s *RPCServer) AddUser(ctx context.Context, res *pb.AddRequest) (*pb.AddResponse, error){
	fmt.Println("controller")
	user := &models.User{
		Name:     res.Name,
		Email:    res.Email,
		Password: res.Password,
		Phone:    res.Phone,
		DOB:      res.Dob,
	}
	fmt.Println("user", user)
	result, err := UserService.AddUser(user)
	if err != nil{
		return &pb.AddResponse{
			Message: "failed",
		}, err
	}
	return &pb.AddResponse{
		Message: result,
	}, nil
	// return nil, nil
}

func (s *RPCServer) AddRole(ctx context.Context, res *pb.RoleRequest) (*pb.AddResponse, error){
	role := &models.Role{
		Name : res.Name,
		Role : res.Role,
	}
	result, err := UserService.AddRole(role)
	if err != nil{
		return &pb.AddResponse{
			Message: "failed",
		}, err
	}
	return &pb.AddResponse{
		Message: result,
	}, nil
}