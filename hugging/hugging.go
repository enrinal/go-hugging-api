package hugging

import "github.com/redis/go-redis/v9"

type Hugging struct {
	Token string
	Rdb   *redis.Client
}
