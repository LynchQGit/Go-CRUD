package main

import (
	"github.com/LynchQGit/Go-CRUD/myweb/internal/pb"
	"github.com/LynchQGit/Go-CRUD/myweb/internal/server"
	"github.com/LynchQGit/Go-CRUD/myweb/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	go server.RunGinServer()  // 启动gin服务

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &service.GreeterService{})
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}