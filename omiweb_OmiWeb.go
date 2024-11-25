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
	SourcePath           string
	IndexPath            string
}

func newOmiWeb(opts *redis.Options, serverName, address string) *OmiWeb {
	return &OmiWeb{
		serverName:           serverName,
		address:              address,
		ServerRegister:       omiserd.NewClient(opts, omiserd.Server).NewRegister(serverName, address),
		opts:                 opts,
		PathProxy:            omiproxy.NewClient(opts).NewProxy(serverName, address, omiproxy.PathMode),
		EnableServerRegister: true,
		SourcePath:           sourcePath,
		IndexPath:            indexPath,
	}
}

func (omiWeb *OmiWeb) GenerateTemplate() {
	copyEmbeddedFiles(omiWeb.SourcePath)
}
func (omiWeb *OmiWeb) EmbedSource(embeddedSource embed.FS) {
	omiWeb.embeddedSource = embeddedSource
	omiWeb.embedModel = true
}

func (omiWeb *OmiWeb) SetCache(cacheDir string, maxSize int) {
	omiWeb.cache = omicafe.NewFileCache(cacheDir, maxSize)
}

func (omiWeb *OmiWeb) handleFunc(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path
	if r.URL.Path == "/" {
		filePath = omiWeb.IndexPath
	}
	filePath = omiWeb.SourcePath + filePath
	var data []byte
	if omiWeb.embedModel {
		data, _ = omiWeb.embeddedSource.ReadFile(filePath)
	} else {
		data, _ = os.ReadFile(filePath)
	}
	w.Write(data)
}

func (omiWeb *OmiWeb) AddHandleFunc(pattern string, handFunc func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(pattern, handFunc)
}

func (omiWeb *OmiWeb) Start(weight int) {
	if omiWeb.EnableServerRegister {
		omiWeb.ServerRegister.RegisterAndServe(weight, func(port string) {})
	}
	omiWeb.PathProxy.SetFailCallback(func(w http.ResponseWriter, r *http.Request) {
		omiWeb.handleFunc(w, r)
	})
	omiWeb.PathProxy.Start(omiproxy.Http, weight, "", "")
}
