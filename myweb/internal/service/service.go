package service

import (
	"context"
	"github.com/LynchQGit/Go-CRUD/myweb/internal/pb"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
