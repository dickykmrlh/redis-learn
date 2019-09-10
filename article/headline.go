package article

import (
	"strconv"
)

func (a *Article) AddHeadline(id int, headline string) {
	key := "article:" + strconv.Itoa(id) + ":headline"
	a.redis.Set(key, headline, 0)
}
