package controllers

import (
	"context"
	"log"

	"github.com/abdulmanafc2001/microservice-api-usercrud/client/models"
	pb "github.com/abdulmanafc2001/microservice-api-usercrud/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":6000"

func GetClient() pb.UserServiceClient {
	con, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Dialing error: %v", err)
	}

	cli := pb.NewUserServiceClient(con)
	return cli
}

func CreateUser(c *gin.Context) {
	var input models.User

	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	cli := GetClient()

	req := &pb.CreateRequest{
		UserName: input.UserName,
		Password: input.Password,
		FullName: input.FullName,
		Phone:    input.Phone,
		Gender:   input.Gender,
	}

	res, err := cli.CreateUser(context.Background(), req)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to get response",
		})
	}

	c.JSON(200, gin.H{
		"success": res.Success.SuccessMessage,
	})

}

func GetAllUsers(c *gin.Context) {

	cli := GetClient()

	res, err := cli.ReadUser(context.Background(), &pb.NoParam{})
	if err != nil {
		c.JSON(400, gin.H{
			"error": "failed to find users",
		})
		return
	}
	c.JSON(200, gin.H{
		"users": res.Userlist,
	})
}
