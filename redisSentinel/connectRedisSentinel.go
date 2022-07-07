package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ExampleClient() {
	fmt.Println("-----start connec redis-------")
	rdb := redis.NewFailoverClusterClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"0.0.0.0:26376", "0.0.0.0:26375", "0.0.0.0:26379"},
		DB:            0, // use default DB
		Password:      "CkkUdLPMT",
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
func main() {
	fmt.Println("-- Start connect --")
	ExampleClient()
}
