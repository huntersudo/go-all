package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

// 发布订阅（pub/sub）是一种消息通信模式
// 发布订阅（pub/sub）是一种消息通信模式，主要的目的是解除消息发布者和消息订阅者之间的耦合，Redis作为一个pub/sub的server，在订阅者和发布者之间起到了消息路由的功能。
//订阅者可以通过subscribe和psubscribe命令向Redis server 订阅自己感兴趣的消息类型，redis将信息类型称为通道（channel）。
//当发布者通过publish命令向Redis server发送特定类型的信息时，订阅该信息类型的全部client都会收到此消息。
// https://www.cnblogs.com/aaron911/p/7862394.html
// https://www.cnblogs.com/jimisun/p/10045772.html

var client2 *redis.Client

func init() {
	client2 = redis.NewClient(&redis.Options{
		Addr:         "192.168.88.132:6379",
		Password:   "iam59!z$",
		DB:       0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
}

func ExamplePubSub( client2 *redis.Client) {
	// 订阅给定的一个或多个频道的信息。
	pubsub := client2.Subscribe("mychannel1")
	// 订阅一个或多个符合给定模式的频道
	//pubsub := client2.PSubscribe("mychannel1")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// Publish a message.
	// 将信息发送到指定的频道。
	err = client2.Publish("mychannel1", "hello").Err()
	if err != nil {
		panic(err)
	}

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})

	// Consume messages.
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}

	// Output: mychannel1 hello
}

func ExamplePubSub_Receive( client2 *redis.Client) {
	pubsub := client2.Subscribe("mychannel2")
	defer pubsub.Close()

	for i := 0; i < 2; i++ {
		// ReceiveTimeout is a low level API. Use ReceiveMessage instead.
		msgi, err := pubsub.ReceiveTimeout(time.Second)
		if err != nil {
			break
		}

		switch msg := msgi.(type) {
		case *redis.Subscription:
			fmt.Println("subscribed to", msg.Channel)

			_, err := client2.Publish("mychannel2", "hello").Result()
			if err != nil {
				panic(err)
			}
		case *redis.Message:
			fmt.Println("received", msg.Payload, "from", msg.Channel)
		default:
			panic("unreached")
		}
	}

	// sent message to 1 rdb
	// received hello from mychannel2
}

func main() {





}