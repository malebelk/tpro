package regular

import (
	"context"
	"github.com/elazarl/goproxy"
	"log"
	"net"
	"net/http"
	"time"
)

type Proxy struct {
	endpoint string
	proxy    *goproxy.ProxyHttpServer
}

func NewProxy(endpoint string, dns string) *Proxy {
	proxy := &Proxy{endpoint: endpoint}

	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, dns+":53")
		},
	}

	net.DefaultResolver = r

	proxy.proxy = goproxy.NewProxyHttpServer()
	proxy.proxy.Verbose = true

	log.Printf("Listening regular proxy on %s", proxy.endpoint)
	log.Fatal(http.ListenAndServe(proxy.endpoint, proxy.proxy))

	return proxy
}
