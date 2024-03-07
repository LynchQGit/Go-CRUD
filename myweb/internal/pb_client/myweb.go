package main


import (
	"context"
	"log"
	"google.golang.org/grpc"
	pb "github.com/LynchQGit/Go-CRUD/myweb/internal/pb" // 确保导入路径与你的项目匹配
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 调用服务
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}