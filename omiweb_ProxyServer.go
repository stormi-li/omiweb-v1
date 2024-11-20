package omiweb

import (
	"log"
	"net/http"

	"github.com/stormi-li/omicafe-v1"
	"github.com/stormi-li/omiserd-v1"
)

type ProxyServer struct {
	router      *router
	webRegister *omiserd.Register
	serverName  string
	address     string
	cache       *omicafe.FileCache
}

func (proxyServer *ProxyServer) handleFunc(w http.ResponseWriter, r *http.Request) {
	domainNameResolution(r, proxyServer.router)
	httpProxy(w, r, proxyServer.cache)
	websocketProxy(w, r)
}

func (proxyServer *ProxyServer) SetCache(cacheDir string, maxSize int) {
	proxyServer.cache = omicafe.NewFileCache(cacheDir, maxSize)
}

func (proxyServer *ProxyServer) StartHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxyServer.handleFunc(w, r)
	})

	proxyServer.webRegister.RegisterAndListen(1, func(port string) {
		log.Println("omi web server: " + proxyServer.serverName + " is running on http://" + proxyServer.address)
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	})
}

func (proxyServer *ProxyServer) StartHttps(cert, key string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxyServer.handleFunc(w, r)
	})

	proxyServer.webRegister.RegisterAndListen(1, func(port string) {
		log.Println("omi web server: " + proxyServer.serverName + " is running on https://" + proxyServer.address)
		err := http.ListenAndServeTLS(port, cert, key, nil)
		if err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	})
}
