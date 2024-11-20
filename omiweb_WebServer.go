package omiweb

import (
	"embed"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/stormi-li/omicafe-v1"
	"github.com/stormi-li/omiserd-v1"
)

type WebServer struct {
	router         *router
	webRegister    *omiserd.Register
	serverName     string
	address        string
	embeddedSource embed.FS
	embedModel     bool
	cache          *omicafe.FileCache
}

func (webServer *WebServer) EmbedSource(embeddedSource embed.FS) {
	webServer.embeddedSource = embeddedSource
	webServer.embedModel = true
}

func (webServer *WebServer) SetCache(cacheDir string, maxSize int) {
	var err error
	webServer.cache = omicafe.NewFileCache(cacheDir, maxSize)
	if err != nil {
		panic(err)
	}
}

func (webServer *WebServer) handleFunc(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) > 0 && webServer.router.Has(parts[1]) {
		pathRequestResolution(r, webServer.router)
		httpProxy(w, r, webServer.cache)
		websocketProxy(w, r)
		return
	}

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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		webServer.handleFunc(w, r)
	})
	webServer.webRegister.RegisterAndListen(weight, func(port string) {
		log.Println("omi web server: " + webServer.serverName + " is running on http://" + webServer.address)
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	})
}
