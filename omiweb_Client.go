package omiweb

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiserd-v1"
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
		router:      newRouter(c.opts, omiserd.Server),
		webRegister: omiserd.NewClient(c.opts, omiserd.Web).NewRegister(serverName, address),
	}
}

func (c *Client) NewProxyServer(serverName, address string) *ProxyServer {
	return &ProxyServer{
		router:      newRouter(c.opts, omiserd.Web),
		serverName:  serverName,
		address:     address,
		webRegister: omiserd.NewClient(c.opts, omiserd.Web).NewRegister(serverName, address),
	}
}
