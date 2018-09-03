package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7001", ":7002", ":7003", ":7004", ":7005", ":7006"},
	})
	client.Ping()

	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
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
