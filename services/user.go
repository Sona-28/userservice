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
	usercoll *mongo.Collection
	rolecoll *mongo.Collection
}

func NewUserService(ctx context.Context, usercoll *mongo.Collection, rolecoll *mongo.Collection) *UserService {
	return &UserService{ctx, usercoll, rolecoll}
}

func (u *UserService) AddUser(user *models.User) (string, error) {
	fmt.Println("service")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)
	_, err = u.usercoll.InsertOne(u.ctx, &user)
	if err != nil {
		return "", err
	}
	return "User added succesfully", nil
}

func (u *UserService) UpdateRole(role *models.UpdateRole) (string, error) {
	filter := bson.D{{Key: "name", Value: role.Name},{Key: "status", Value: "enabled"}}
	var result *models.User
	err := u.usercoll.FindOne(u.ctx, filter).Decode(&result)
	if err != nil {
		return "", err
	}
	_, err = u.usercoll.UpdateOne(u.ctx, filter, bson.D{{Key: "$set", Value: bson.D{{Key: "role", Value: role.Role}}}})
	if err != nil {
		return "", err
	}
	return "Role updated succesfully", nil
}

func (u *UserService) ListFeatures(list *models.Role) (*models.Role, error){
	filter := bson.D{{Key: "role", Value: list.Role}}
	var result *models.Role
	err := u.rolecoll.FindOne(u.ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserService) DisableUser(list *models.User) (string, error){
	filter := bson.D{{Key: "name", Value: list.Name}}
	_,err := u.usercoll.UpdateOne(u.ctx, filter, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "disabled"}}}})
	if err != nil {
		return "", err
	}
	return "User disabled succesfully", nil

}

func (u *UserService) EnableUser(list *models.User) (string, error){
	filter := bson.D{{Key:"name", Value: list.Name}}
	_,err := u.usercoll.UpdateOne(u.ctx, filter, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "enabled"}}}})
	if err != nil{
		return "", err
	}
	return "User enabled succesfully", nil
}

func (u *UserService) AssociateRole(list *models.RoleRequest) (string, error){
	filter := bson.D{{Key:"name", Value: list.Name}}
	_, err := u.usercoll.UpdateOne(u.ctx, filter, bson.D{{Key: "$push", Value: bson.D{{Key: "role",Value: list.Role}}}})
	if err != nil{
		return "", err
	}
	return "Role added succesfully", nil
}