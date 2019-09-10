package article

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
)

type Vote struct {
	redis *redis.Client
}

func InitVote(client *redis.Client) *Vote {
	return &Vote{
		redis: client,
	}
}

func (a *Vote) UpVote(id int) {
	key := "article:" + strconv.Itoa(id) + ":votes"
	a.redis.Incr(key)
}

func (a *Vote) DownVote(id int) {
	key := "article:" + strconv.Itoa(id) + ":votes"
	a.redis.Decr(key)
}

func (a *Vote) ShowResults(id int) {
	headlineKey := "article:" + strconv.Itoa(id) + ":headline"
	voteKey := "article:" + strconv.Itoa(id) + ":votes"
	sc := a.redis.MGet(headlineKey, voteKey)
	result, _ := sc.Result()
	fmt.Printf(`The Article: "%s" has %s votes`, result[0], result[1])
	fmt.Println()
}
