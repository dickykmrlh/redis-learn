package article

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

type Article struct {
	redis *redis.Client
}

func InitVote(client *redis.Client) *Article {
	return &Article{
		redis: client,
	}
}

func (a *Article) UpVote(id int) {
	key := "article:" + strconv.Itoa(id) + ":votes"
	a.redis.Incr(key)
}

func (a *Article) DownVote(id int) {
	key := "article:" + strconv.Itoa(id) + ":votes"
	a.redis.Decr(key)
}

func (a *Article) ShowResults(id int) {
	headlineKey := "article:" + strconv.Itoa(id) + ":headline"
	voteKey := "article:" + strconv.Itoa(id) + ":votes"
	sc := a.redis.MGet(headlineKey, voteKey)
	result, _ := sc.Result()
	fmt.Printf(`The Article: "%s" has %s votes`, result[0], result[1])
	fmt.Println()
}
