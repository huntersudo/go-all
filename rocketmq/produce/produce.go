package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{"192.168.88.132:9876"}),
		//producer.WithNsResolver(primitive.NewPassthroughResolver(endPoint)),
		producer.WithRetry(2),
		producer.WithGroupName("GID_xxxxxx"),
	)
	if err!=nil{
		fmt.Errorf("produce init error:",err)
	}

	err1 := p.Start()
	if err1!=nil{
		fmt.Errorf("produce start error:",err1)
	}
	result, err := p.SendSync(context.Background(), &primitive.Message{
		Topic: "test",
		Body:  []byte("Hello RocketMQ Go Client!"),
	})

	fmt.Println("produce: ",result)

	//or send message with oneway
	//err := p.SendOneWay(context.Background(), &primitive.Message{
	//	Topic: "test",
	//	Body:  []byte("Hello RocketMQ Go Client!"),
	//})
}