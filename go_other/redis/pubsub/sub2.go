package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)
var client_sub2 *redis.Client

func main() {
	client_sub2 = redis.NewClient(&redis.Options{
		Addr:         "192.168.88.132:6379",
		Password:   "iam59!z$",
		DB:       0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	pubsub2(client_sub2)
}
func pubsub2( client2 *redis.Client) {
	pubsub := client2.Subscribe("mychannel2")
	defer pubsub.Close()

	//for i := 0; i < 200; i++ {

		for {
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
