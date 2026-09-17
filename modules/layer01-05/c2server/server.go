package c2server

import (
	"time"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (e *Server) Run() ([]string, error) {
	results := make([]string, 0, 1)
	results = append(results, "server:done")
	return results, nil
}

func (e *Server) Name() string { return "Server" }
func (e *Server) Timestamp() time.Time { return time.Now() }
