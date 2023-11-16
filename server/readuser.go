package main

import (
	"context"
	"log"

	"github.com/abdulmanafc2001/microservice-api-usercrud/client/models"
	pb "github.com/abdulmanafc2001/microservice-api-usercrud/proto"
	db "github.com/abdulmanafc2001/microservice-api-usercrud/server/database"
)

func (s *UserServer) ReadUser(ctx context.Context, req *pb.NoParam) (*pb.UserList, error) {
	var userlist []models.User
	if err := db.ConnectToDB().Find(&userlist).Error; err != nil {
		log.Println("hkdsk", userlist)
		return nil, err
	}
	var Users []*pb.User
	for _, v := range userlist {
		Users = append(Users, &pb.User{
			Id:       int32(v.Id),
			UserName: v.UserName,
			Password: v.Password,
			FullName: v.FullName,
			Phone:    v.Phone,
			Gender:   v.Gender,
		})
	}
	log.Println(Users)
	return &pb.UserList{Userlist: Users}, nil

}
