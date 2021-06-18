package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
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

func main() {
	err := client.Set("key-1", "value-1", 0).Err()
	err = client.Set("key-2", "value-2", 0).Err()
	if err != nil {
		panic(err)
	}

	val1, err := client.Get("key-1").Result()
	val2, err := client.Get("key-2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key-1:  ", val1)
	fmt.Println("key-2:  ", val2)

	val3, err := client.Get("missing_key").Result()
	if err == redis.Nil {
		fmt.Println("missing_key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("missing_key", val3)
	}

}


