package api

import (
    "log"
    "net/http"
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
    log.Println("API server is running on :9000")
    http.ListenAndServe(":9000", nil)
}
