package main

import (
	"context"
	"errors"

	pb "github.com/abdulmanafc2001/microservice-api-usercrud/proto"
	db "github.com/abdulmanafc2001/microservice-api-usercrud/server/database"
	"github.com/abdulmanafc2001/microservice-api-usercrud/server/models"
)

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	userId := req.GetUserId()

	if row := db.ConnectToDB().Where("id = ?", userId).Delete(&models.User{}).RowsAffected; row != 1 {
		return nil, errors.New("id does not exist")
	}
	dres := &pb.DeleteResponse{
		Message: &pb.SuccessMessage{
			SuccessMessage: "successfully deleted",
		},
	}
	return dres, nil
}
