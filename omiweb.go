package omiweb

import (
	"github.com/go-redis/redis/v8"
)

func NewClient(opts *redis.Options) *Client {
	return &Client{
		opts: opts,
	}
}

func DisableLog() {
	log_cache = false
}
