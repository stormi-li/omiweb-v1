package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiweb-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	c := omiweb.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	c.GenerateTemplate()
	ws := c.NewOmiWeb("test8084", "118.25.196.166:8084")
	ws.Start(1)
}
