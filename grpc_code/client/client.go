package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "grpc_code/proto"
)

const (
	PORT = "9001"
)

func main() {
	//连接grpc server
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()

	//创建 SearchService 的客户端对象
	client := pb.NewSearchServiceClient(conn)

	//发送 RPC 请求，等待同步响应，得到回调后返回响应结果
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp:    %s", resp.GetResponse())
}
