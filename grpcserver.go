package main

import (
	"context"

	"github.com/9d4/inothing-broker-service/pb/userpb"
)

type GrpcServer struct{
	userpb.UnimplementedUserServiceServer
}

func (s *GrpcServer) NewUser(ctx context.Context, in *userpb.NewUserRequest) (*userpb.NewUserResponse, error) {
	return &userpb.NewUserResponse{
		Success: true,
		Message: "Holaa",
	}, nil
}
