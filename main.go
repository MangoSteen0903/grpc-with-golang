package main

import (
	"context"
	"log"
	"net"

	"github.com/MangoSteen0903/grpc-go-tutorial/data"
	pb "github.com/MangoSteen0903/grpc-go-tutorial/user"
	"google.golang.org/grpc"
)

const portNum = "9000"

type userServer struct {
	pb.UserServiceServer
}

func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	userId := req.UserId

	var resMessage *pb.UserMessage

	for _, user := range data.UserData {
		if user.UserId == userId {
			resMessage = user
		}
	}
	return &pb.GetUserResponse{
		UserMessage: resMessage,
	}, nil
}
func main() {

	lis, err := net.Listen("tcp", ":"+portNum)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServer{})
	log.Printf("Start gRPC server on %s port", portNum)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to Serve: %s", err)
	}
}
