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

	article := article.Init(client)
	article1ID := 12345
	article2ID := 10001
	article3ID := 60056
	// set article headline
	article.AddHeadline(article1ID, "Google Wants to Turn Your Clothes Into a Computer")
	article.AddHeadline(article2ID, "For Millennials, the End of the TV Viewing Party")
	article.AddHeadline(article3ID, "Alicia Vikander, Who Portrayed Denmark's Queen, Is Screen Royalty")

	// vote article
	article.UpVote(article1ID)   // article:12345 has 1 vote
	article.UpVote(article1ID)   // article:12345 has 2 vote
	article.UpVote(article1ID)   // article:12345 has 3 vote
	article.UpVote(article2ID)   // article:10001 has 1 vote
	article.UpVote(article2ID)   // article:10001 has 2 vote
	article.DownVote(article2ID) // article:10001 has 1 vote
	article.UpVote(article3ID)   // article:60056 has 1 vote

	// show result
	article.ShowResults(article1ID)
	article.ShowResults(article2ID)
	article.ShowResults(article3ID)
}
