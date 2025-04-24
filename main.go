package main

import (
	"github.com/joho/godotenv"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	endpoint := os.Getenv("ENDPOINT")
	target := os.Getenv("TARGET")

	ln, err := net.Listen("tcp", endpoint)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening on %s", endpoint)

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		log.Printf("Get connection from %s\n", conn.RemoteAddr().String())
		go handleRequest(conn, target)
	}
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
