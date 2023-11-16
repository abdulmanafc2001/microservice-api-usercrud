package controllers

import (
	"context"
	"strconv"

	pb "github.com/abdulmanafc2001/microservice-api-usercrud/proto"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cli := GetClient()
	dlReq := &pb.DeleteRequest{
		UserId: int32(id),
	}
	res, err := cli.DeleteUser(context.Background(), dlReq)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": res.Message.SuccessMessage,
	})
}
