# Go-CRUD
This is my demo of learning the Golang CURD.

## 项目结构
myweb/
├── bin/           # 编译后的二进制文件
├── cmd/           # 主程序入口
│   └── server/    # 服务器启动程序
├── internal/      # 内部包
    ├── pb/        # Protocol Buffers 定义和生成的代码
    ├── pb_client/ # gRPC 客户端示例
    ├── server/    # HTTP 服务器实现
    └── service/   # 业务服务实现

## 开发环境配置

### 必要工具安装
```bash
# 安装 protoc 编译器
brew install protobuf

# 安装 Go 的 protoc 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 开发流程
### 1. 定义服务接口
在 `myweb/internal/pb/myweb.proto` 中定义新的服务接口：
```protobuf
syntax = "proto3";

package myweb;
option go_package = "myweb/internal/pb";

service YourService {
    rpc YourMethod (YourRequest) returns (YourResponse) {}
}

message YourRequest {
    // 定义请求字段
}

message YourResponse {
    // 定义响应字段
}
```

### 2. 生成 Protocol Buffers 代码
```bash
# 在项目根目录下执行
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    myweb/internal/pb/myweb.proto
```

### 3. 实现服务
在 `myweb/internal/service` 目录下创建服务实现：
```go
package service

import (
    "context"
    "github.com/LynchQGit/Go-CRUD/myweb/internal/pb"
    "github.com/sirupsen/logrus"
)

type YourService struct {
    pb.UnimplementedYourServiceServer
}

func (s *YourService) YourMethod(ctx context.Context, req *pb.YourRequest) (*pb.YourResponse, error) {
    logrus.Infof("Received request: %v", req)
    // 实现业务逻辑
    return &pb.YourResponse{}, nil
}
```

### 4. 注册服务
在 `cmd/server/main.go` 中注册新服务：
```go
s := grpc.NewServer()
pb.RegisterYourServiceServer(s, &service.YourService{})
```

## 项目规范
### 目录结构规范
- cmd/ : 存放主程序入口
- internal/ : 存放内部包，对外不可见
- pb/ : 存放 protobuf 相关代码
- service/ : 存放具体的服务实现
- server/ : 存放 HTTP 服务相关代码

### 代码规范
1. 所有服务实现都应该添加适当的日志记录
2. 使用 logrus 进行日志记录
3. 遵循 Go 标准项目布局
4. 使用 gofmt 格式化代码

### Git 提交规范
提交信息格式：

```plaintext
<type>: <description>

[optional body]
 ```

type 类型：

- feat: 新功能
- fix: 修复
- docs: 文档更新
- style: 代码格式
- refactor: 重构
- test: 测试
- chore: 构建过程或辅助工具的变动

## 服务端口
- gRPC 服务：50051
- HTTP 服务：8080

## 测试
1. 使用 pb_client 目录下的测试客户端进行 gRPC 测试
2. 使用 curl 或 Postman 测试 HTTP 接口

## 部署
待补充...