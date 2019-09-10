package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("my_key", "Hello World using go-redis and Redis", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("my_key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("my_key", val)
}
