package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omiweb-v1"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	c := omiweb.NewClient(&redis.Options{Addr: redisAddr, Password: password})
	ws := c.NewOmiWeb("test8085", "118.25.196.166:8085")
	ws.AddHandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})
	ws.Start(1)
}
