package omiweb

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiproxy-v1"
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
		serverName:     serverName,
		address:        address,
		ServerRegister: omiserd.NewClient(c.opts, omiconst.Server).NewRegister(serverName, address),
		opts:           c.opts,
		PathProxy:      omiproxy.NewClient(c.opts).NewProxy(serverName, address, omiproxy.PathMode),
	}
}
