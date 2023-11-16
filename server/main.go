package main

import (
	"log"
	"net"

	pb "github.com/abdulmanafc2001/microservice-api-usercrud/proto"
	db "github.com/abdulmanafc2001/microservice-api-usercrud/server/database"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const port = ":6000"

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func init() {
	log.Println(godotenv.Load(".env"))
	db.ConnectToDB()
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("could not connect with %s and got error %v", port, err)
	}

	grpcServer := grpc.NewServer()

	log.Printf("Server Started at %v", lis.Addr())

	pb.RegisterUserServiceServer(grpcServer, &UserServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Server serving error %v", err)
	}

}
