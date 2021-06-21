package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)


var client_pub *redis.Client

func main() {
	client_pub = redis.NewClient(&redis.Options{
		Addr:         "192.168.88.132:6379",
		Password:   "iam59!z$",
		DB:       0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	pub(client_pub)

}
func pub( client2 *redis.Client) {
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
	err = client2.Publish("mychannel1", "hello--001").Err()
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
