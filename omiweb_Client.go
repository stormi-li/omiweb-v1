package omiweb

import (
	"github.com/go-redis/redis/v8"
)

type Client struct {
	opts *redis.Options
}



func (c *Client) NewOmiWeb(serverName, address string) *OmiWeb {
	return newOmiWeb(c.opts, serverName, address)
}
