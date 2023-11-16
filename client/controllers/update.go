package controllers

import (
	"context"
	"strconv"

	"github.com/abdulmanafc2001/microservice-api-usercrud/client/models"
	pb "github.com/abdulmanafc2001/microservice-api-usercrud/proto"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	var input models.User
	if err := c.Bind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to bind data",
		})
		return
	}

	cli := GetClient()

	id, _ := strconv.Atoi(c.Param("id"))

	upReq := &pb.UpdateRequest{
		UserId:   int32(id),
		UserName: input.UserName,
		Password: input.Password,
		FullName: input.FullName,
		Phone:    input.Phone,
		Gender:   input.Gender,
	}

	res, err := cli.UpdateUser(context.Background(), upReq)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to get response",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": res.Message.SuccessMessage,
	})
}
