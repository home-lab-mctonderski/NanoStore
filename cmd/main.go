package main

import (
    "github.com/McTonderski/storage-service/internal/api"
    "log"
)

func main() {
	log.Println("Starting storage-service...")
	server := api.NewServer()
	server.Start()
}
