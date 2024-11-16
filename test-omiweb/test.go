package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiweb-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	omiweb := omiweb.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	omiweb.GenerateTemplate()
	ws := omiweb.NewWebServer("118.25.196.166", "118.25.196.166:7073")
	ws.Start(1)
}
