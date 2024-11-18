package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiweb-v1"
)

func main() {
	omiweb := omiweb.NewClient(&redis.Options{Addr: "localhost:6379"})
	omiweb.GenerateTemplate()
	ws := omiweb.NewWebServer("localhost", "localhost:8080")
	ws.Start(1)
}
