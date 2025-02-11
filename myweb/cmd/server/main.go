package main

import (
	"github.com/LynchQGit/Go-CRUD/myweb/internal/pb"
	"github.com/LynchQGit/Go-CRUD/myweb/internal/server"
	"github.com/LynchQGit/Go-CRUD/myweb/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	go server.RunGinServer() // 启动gin服务

	listen, err := net.Listen("tcp", ":50051") // 监听端口
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()                                  // 创建grpc服务器
	pb.RegisterGreeterServer(s, &service.GreeterService{}) // 注册服务
	err = s.Serve(listen)                                  // 启动grpc服务器
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
