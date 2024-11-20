package omiweb

import (
	"embed"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/omicafe-v1"
	"github.com/stormi-li/omiproxy-v1"
	"github.com/stormi-li/omiserd-v1"
)

type WebServer struct {
	webRegister    *omiserd.Register
	serverName     string
	address        string
	embeddedSource embed.FS
	embedModel     bool
	cache          *omicafe.FileCache
	opts           *redis.Options
	reverseProxy   *omiproxy.ReverseProxy
}

func (webServer *WebServer) EmbedSource(embeddedSource embed.FS) {
	webServer.embeddedSource = embeddedSource
	webServer.embedModel = true
}

func (webServer *WebServer) SetCache(cacheDir string, maxSize int) {
	webServer.cache = omicafe.NewFileCache(cacheDir, maxSize)
}

func (webServer *WebServer) handleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	if r.URL.Path == "/" {
		filePath = index_path
	}
	filePath = target_path + filePath
	var data []byte
	if webServer.embedModel {
		data, _ = webServer.embeddedSource.ReadFile(filePath)
	} else {
		data, _ = os.ReadFile(filePath)
	}
	w.Write(data)
}

func (webServer *WebServer) Start(weight int) {
	webServer.reverseProxy = omiproxy.NewClient(webServer.opts).NewReverseProxy(webServer.serverName, webServer.address)
	webServer.reverseProxy.SetFailCallback(func(w http.ResponseWriter, r *http.Request) {
		webServer.handleFunc(w, r)
	})
	webServer.reverseProxy.Start(weight)
}
