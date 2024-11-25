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

type OmiWeb struct {
	serverName           string
	address              string
	embeddedSource       embed.FS
	embedModel           bool
	cache                *omicafe.FileCache
	opts                 *redis.Options
	PathProxy            *omiproxy.OmiProxy
	ServerRegister       *omiserd.Register
	EnableServerRegister bool
}

func newOmiWeb(opts *redis.Options, serverName, address string) *OmiWeb {
	return &OmiWeb{
		serverName:           serverName,
		address:              address,
		ServerRegister:       omiserd.NewClient(opts, omiserd.Server).NewRegister(serverName, address),
		opts:                 opts,
		PathProxy:            omiproxy.NewClient(opts).NewProxy(serverName, address, omiproxy.PathMode),
		EnableServerRegister: true,
	}
}
func (webServer *OmiWeb) EmbedSource(embeddedSource embed.FS) {
	webServer.embeddedSource = embeddedSource
	webServer.embedModel = true
}

func (webServer *OmiWeb) SetCache(cacheDir string, maxSize int) {
	webServer.cache = omicafe.NewFileCache(cacheDir, maxSize)
}

func (webServer *OmiWeb) handleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	if r.URL.Path == "/" {
		filePath = IndexPath
	}
	filePath = StaticPath + filePath
	var data []byte
	if webServer.embedModel {
		data, _ = webServer.embeddedSource.ReadFile(filePath)
	} else {
		data, _ = os.ReadFile(filePath)
	}
	w.Write(data)
}

func (webServer *OmiWeb) AddHandleFunc(pattern string, handFunc func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handFunc)
}

func (webServer *OmiWeb) Start(weight int) {
	if webServer.EnableServerRegister {
		webServer.ServerRegister.RegisterAndServe(weight, func(port string) {})
	}
	webServer.PathProxy.SetFailCallback(func(w http.ResponseWriter, r *http.Request) {
		webServer.handleFunc(w, r)
	})
	webServer.PathProxy.Start(omiproxy.Http, weight, "", "")
}
