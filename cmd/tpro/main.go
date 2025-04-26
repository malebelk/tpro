package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"tpro/internal/regular"
	"tpro/internal/reverse"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env.example file")
	}

	proxyType := os.Getenv("TYPE")
	endpoint := os.Getenv("ENDPOINT")
	target := os.Getenv("TARGET")
	dns := os.Getenv("DNS")

	switch proxyType {
	case "reverse":
		reverse.NewProxy(endpoint, target)
	case "regular":
		regular.NewProxy(endpoint, dns)
	default:
		regular.NewProxy(endpoint, dns)
	}
}
