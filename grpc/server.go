package grpc

import "github.com/9d4/inothing-broker-service/pb/userpb"

type GrpcServer struct{
	userpb.UnimplementedUserServiceServer
}
