package main

import (
	"context"
	"log"
	"net"

	pb "grpc_code/proto"

	"google.golang.org/grpc"
)

type SearchService struct{}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	return &pb.SearchResponse{Response: r.GetRequest() + " Server"}, nil
}

const (
	PORT = "9001"
)

func main() {
	server := grpc.NewServer() //创建 gRPC Server对象

	//将 SearchService（其包含需要被调用的服务端接口）注册到gRPC Server 的内部注册中心
	//这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行逻辑处理
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", ":"+PORT) //创建 Listen，监听 TCP 端口
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	//gRPC Server开始 lis.Accept,直到 Stop 或 GracefulStop
	server.Serve(lis)
}
