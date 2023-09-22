package main

import (
	"context"
	"net/http"
	pb "sample/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main(){
	r := gin.Default()
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)
	r.POST("/add",func(c *gin.Context){
		var request pb.AddRequest
		if err := c.ShouldBindJSON(&request); err!=nil{
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		res, err := client.AddUser(context.TODO(), &pb.AddRequest{
			Name:     request.Name,
			Email:    request.Email,
			Password: request.Password,
			Dob:      request.Dob,
			Phone:    request.Phone,
		})
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	})
	r.POST(("/role"), func(c *gin.Context){
		var request pb.RoleRequest
		if err := c.ShouldBindJSON(&request); err!=nil{
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		res, err := client.AddRole(context.TODO(), &pb.RoleRequest{
			Name: request.Name,
			Role: request.Role,
		})
		if err != nil{
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})		
	})
	r.Run(":2000")
}