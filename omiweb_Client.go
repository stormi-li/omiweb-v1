package omiweb

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiserd-v1"
	omiconst "github.com/stormi-li/omiserd-v1/omiserd_const"
)

type Client struct {
	opts *redis.Options
}

func (c *Client) GenerateTemplate() {
	copyEmbeddedFiles()
}

func (c *Client) NewWebServer(serverName, address string) *WebServer {
	return &WebServer{
		serverName:  serverName,
		address:     address,
		webRegister: omiserd.NewClient(c.opts, omiconst.Web).NewRegister(serverName, address),
		opts:        c.opts,
	}
}
