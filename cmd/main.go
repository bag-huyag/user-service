package main

import (
	"log"

	"github.com/bag-huyag/user-service/pkg/server"
)

func main() {
	if err := server.StartGRPC(); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
