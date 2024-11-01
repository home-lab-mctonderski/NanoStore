// Author: Maciej Tonderski
package main

import (
	"log"

	"github.com/McTonderski/storage-service/internal/api"
)

func main() {
	log.Println("Starting storage-service...")
	server := api.NewServer()
	server.Start()
}
