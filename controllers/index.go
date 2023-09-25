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
		Role:    res.Role,
		Status:   res.Status,
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

func (s *RPCServer) UpdateRole(ctx context.Context, res *pb.UpdateRoleRequest) (*pb.AddResponse, error){
	role := &models.UpdateRole{
		Name : res.Name,
		Role : res.Role,
	}
	result, err := UserService.UpdateRole(role)
	if err != nil{
		return &pb.AddResponse{
			Message: "failed",
		}, err
	}
	return &pb.AddResponse{
		Message: result,
	}, nil
}

func (s *RPCServer) ListFeature(ctx context.Context, res *pb.ListRequest) (*pb.Role, error){
	list := &models.Role{
		Role: res.Role,
	}
	result, err := UserService.ListFeatures(list)
	if err!=nil{
		return nil, err
	}
	listreq := &pb.Role{
		Role:           result.Role,
		Responsibility: result.Responsibility,
		Access:         result.Access,
	}
	return listreq, nil
}

func (s *RPCServer) EnableUser(ctx context.Context, res *pb.UserRequest) (*pb.AddResponse, error){
	list := &models.User{
		Name: res.Name,
	}
	result, err := UserService.EnableUser(list)
	if err!=nil{
		return &pb.AddResponse{
			Message: "failed",
		}, err
	}
	return &pb.AddResponse{
		Message: result,
	}, nil
}

func (s *RPCServer) DisableUser(ctx context.Context, res *pb.UserRequest) (*pb.AddResponse, error){
	list := &models.User{
		Name: res.Name,
	}
	result, err := UserService.DisableUser(list)
	if err!=nil{
		return &pb.AddResponse{
			Message: "failed",
		}, err
	}
	return &pb.AddResponse{
		Message: result,
	}, nil
}

func (s *RPCServer) AssociateRole(ctx context.Context, res *pb.RoleRequest) (*pb.AddResponse, error){
	list := &models.RoleRequest{
		Name: res.Name,
		Role: res.Role,
	}
	result, err := UserService.AssociateRole(list)
	if err!=nil{
		return &pb.AddResponse{
			Message: "failed",
		}, err
	}
	return &pb.AddResponse{
		Message: result,
	}, nil
	
}