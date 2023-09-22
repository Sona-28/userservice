package services

import (
	"context"
	"fmt"
	"sample/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	ctx   context.Context
	mcoll *mongo.Collection
}

func NewUserService(ctx context.Context, mcoll *mongo.Collection) *UserService {
	return &UserService{ctx, mcoll}
}

func (u *UserService) AddUser(user *models.User) (string, error) {
	fmt.Println("service")
	fmt.Println("userserv", user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)
	_, err = u.mcoll.InsertOne(u.ctx, &user)
	if err != nil {
		return "", err
	}
	return "User added succesfully", nil
}

func (u *UserService) AddRole(role *models.Role) (string, error) {
	filter := bson.D{{Key: "name", Value: role.Name}}
	var result models.Role
	err := u.mcoll.FindOne(u.ctx, filter).Decode(&result)
	if err != nil {
		return "", err
	}
	_, err = u.mcoll.UpdateOne(u.ctx, filter, bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: role.Role}}}})
	if err != nil {
		return "", err
	}
	return "Role updated succesfully", nil
}
