package service

import (
	"context"
	"github.com/LynchQGit/Go-CRUD/myweb/internal/pb"
	"github.com/sirupsen/logrus"
)

type GreeterService struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	logrus.Infof("Received SayHello request with name: %s", in.Name)
	reply := &pb.HelloReply{Message: "Hello " + in.Name}
	logrus.Infof("Sending response: %s", reply.Message)
	return reply, nil
}
