// Author: Maciej Tonderski
package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Server struct
type Server struct {
	// add server configurations and dependencies here
}

// NewServer creates a new API server
func NewServer() *Server {
	return &Server{}
}

// Start the API server
func (s *Server) Start() {
	router := gin.Default()

	// Define your routes and handlers here
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	server := &http.Server{
		Addr:              ":9000",
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	log.Println("Starting server on :9000")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
