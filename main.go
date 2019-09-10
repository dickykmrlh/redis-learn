package main

import (
	"fmt"
	"log"

	"github.com/dickyboKa/redis-learn/article"
	"github.com/go-redis/redis"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(pong, err)

	voteSystem := article.InitVote(client)

	voteSystem.UpVote(12345)   // article:12345 has 1 vote
	voteSystem.UpVote(12345)   // article:12345 has 2 vote
	voteSystem.UpVote(12345)   // article:12345 has 3 vote
	voteSystem.UpVote(10001)   // article:10001 has 1 vote
	voteSystem.UpVote(10001)   // article:10001 has 2 vote
	voteSystem.DownVote(10001) // article:10001 has 1 vote
	voteSystem.UpVote(60056)   // article:60056 has 1 vote

	voteSystem.ShowResults(12345)
	voteSystem.ShowResults(10001)
	voteSystem.ShowResults(60056)
}
