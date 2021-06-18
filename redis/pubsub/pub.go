package main

import (
	"github.com/go-redis/redis/v7"
	"time"
)

func main() {
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
	ExamplePubSub(client2)

}
