package main

import (
	"context"

	pb "github.com/abdulmanafc2001/microservice-api-usercrud/proto"
	db "github.com/abdulmanafc2001/microservice-api-usercrud/server/database"
	"github.com/abdulmanafc2001/microservice-api-usercrud/server/models"
)

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	fullName := req.GetFullName()
	userName := req.GetUserName()
	password := req.GetPassword()
	phone := req.GetPhone()
	gender := req.GetGender()

	if err := db.ConnectToDB().Create(&models.User{
		UserName: userName,
		Password: password,
		FullName: fullName,
		Phone:    phone,
		Gender:   gender,
	}).Error; err != nil {
		return nil, err
	}
	res := &pb.CreateResponse{
		Success: &pb.SuccessMessage{
			SuccessMessage: "Successfully created " + userName,
		},
	}

	return res, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	fullName := req.GetFullName()
	userName := req.GetUserName()
	password := req.GetPassword()
	phone := req.GetPhone()
	gender := req.GetGender()
	userid := req.GetUserId()

	if err := db.ConnectToDB().Model(&models.User{}).Where("id=?", userid).Updates(map[string]interface{}{
		"full_name": fullName,
		"user_name": userName,
		"password":  password,
		"phone":     phone,
		"gender":    gender,
	}).Error; err != nil {
		return nil, err
	}
	res := &pb.UpdateResponse{
		Message: &pb.SuccessMessage{
			SuccessMessage: "Successfully updated user",
		},
	}
	return res, nil
}
