package regular

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

type Proxy struct {
	endpoint string
	proxy    *goproxy.ProxyHttpServer
}

func NewProxy(endpoint string) *Proxy {
	proxy := &Proxy{endpoint: endpoint}

	proxy.proxy = goproxy.NewProxyHttpServer()
	proxy.proxy.Verbose = true
	log.Printf("Listening regular proxy on %s", proxy.endpoint)
	log.Fatal(http.ListenAndServe(proxy.endpoint, proxy.proxy))

	return proxy
}
