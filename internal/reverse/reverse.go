package reverse

import (
	"io"
	"log"
	"net"
)

type Proxy struct {
	endpoint string
	target   string
}

func NewProxy(endpoint string, target string) *Proxy {
	proxy := &Proxy{endpoint: endpoint, target: target}

	ln, err := net.Listen("tcp", proxy.endpoint)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening reverse proxy on %s, redirect to %s", proxy.endpoint, proxy.target)

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		log.Printf("Get connection from %s\n", conn.RemoteAddr().String())
		go handleRequest(conn, proxy.target)
	}

	return proxy
}

func handleRequest(conn net.Conn, target string) {
	proxy, err := net.Dial("tcp", target)
	if err != nil {
		panic(err)
	}

	go copyIO(conn, proxy)
	go copyIO(proxy, conn)
}

func copyIO(src, dest net.Conn) {
	defer src.Close()
	defer dest.Close()
	io.Copy(src, dest)
}
