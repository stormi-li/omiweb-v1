package main

import (
	"embed"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiweb-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

//go:embed static/*
var embeddedSource embed.FS

func main() {
	omiweb := omiweb.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	ws := omiweb.NewWebServer("demo.stormili.site", "118.25.196.166:8181")
	ws.Start(1)
}
